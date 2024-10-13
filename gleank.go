package gleank

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/Torbatti/gleank/cmd"
	"github.com/Torbatti/gleank/core"
	"github.com/Torbatti/gleank/utils"
	"github.com/spf13/cobra"
)

const Version = "(untracked)"

var _ core.App = (*Gleank)(nil)

type appWrapper struct {
	core.App
}

type Gleank struct {
	*appWrapper

	devFlag     bool
	dataDirFlag string

	RootCmd *cobra.Command // RootCmd is the main console command
}

type Config struct {
	// optional default values for the console flags
	DefaultDev     bool
	DefaultDataDir string // if not set, it will fallback to "./gk_data"
}

func New() *Gleank {
	_, isUsingGoRun := inspectRuntime()

	return NewWithConfig(Config{
		DefaultDev: isUsingGoRun,
	})
}

func NewWithConfig(config Config) *Gleank {
	// initialize a default data directory based on the executable baseDir
	if config.DefaultDataDir == "" {
		baseDir, _ := inspectRuntime()
		config.DefaultDataDir = filepath.Join(baseDir, "gk_data")
	}

	gk := &Gleank{
		RootCmd: &cobra.Command{
			Use:     filepath.Base(os.Args[0]),
			Short:   "Gleank CLI",
			Version: Version,

			// no need to provide the default cobra completion command
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
		devFlag:     config.DefaultDev,
		dataDirFlag: config.DefaultDataDir,
	}

	// initialize the app instance
	gk.appWrapper = &appWrapper{core.NewBaseApp(core.BaseAppConfig{
		IsDev:   gk.devFlag,
		DataDir: gk.dataDirFlag,
	})}

	return gk
}

func (gk *Gleank) Start() error {
	gk.RootCmd.AddCommand(cmd.NewServeCommand(gk))

	return gk.Execute()
}
func (gk *Gleank) Execute() error {
	if !gk.skipBootstrap() {
		if err := gk.Bootstrap(); err != nil {
			log.Fatal(err)
		}
	}

	done := make(chan bool, 1)

	// listen for interrupt signal to gracefully shutdown the application
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

		done <- true
	}()

	// execute the root command
	go func() {
		// note: leave to the commands to decide whether to print their error
		gk.RootCmd.Execute()

		done <- true
	}()

	<-done

	// TODO : Trigger Cleanups

	return fmt.Errorf("")
}

// skipBootstrap eagerly checks if the app should skip the bootstrap process:
// - already bootstrapped
// - is unknown command
// - is the default help command
// - is the default version command
//
// https://github.com/pocketbase/pocketbase/issues/404
// https://github.com/pocketbase/pocketbase/discussions/1267
func (gk *Gleank) skipBootstrap() bool {
	flags := []string{
		"-h",
		"--help",
		"-v",
		"--version",
	}

	if gk.IsBootstrapped() {
		return true // already bootstrapped
	}

	cmd, _, err := gk.RootCmd.Find(os.Args[1:])
	if err != nil {
		return true // unknown command
	}

	for _, arg := range os.Args {
		if !utils.ExistInSlice(arg, flags) {
			continue
		}

		// ensure that there is no user defined flag with the same name/shorthand
		trimmed := strings.TrimLeft(arg, "-")
		if len(trimmed) > 1 && cmd.Flags().Lookup(trimmed) == nil {
			return true
		}
		if len(trimmed) == 1 && cmd.Flags().ShorthandLookup(trimmed) == nil {
			return true
		}
	}

	return false
}

// inspectRuntime tries to find the base executable directory and how it was run.
func inspectRuntime() (baseDir string, withGoRun bool) {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// probably ran with go run
		withGoRun = true
		baseDir, _ = os.Getwd()
	} else {
		// probably ran with go build
		withGoRun = false
		baseDir = filepath.Dir(os.Args[0])
	}
	return
}
