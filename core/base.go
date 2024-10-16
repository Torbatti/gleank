package core

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Torbatti/gleank/models"
	"github.com/Torbatti/gleank/models/settings"
	sqlz "github.com/Torbatti/gleank/models/sqlc"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var _ App = (*BaseApp)(nil)

// BaseApp implements core.App and defines the base Gleank app structure.
type BaseApp struct {
	// configurable parameters
	isDev   bool
	dataDir string

	// internals
	db         *sql.DB
	db_context context.Context
	store      *redka.DB
	tokenAuth  *jwtauth.JWTAuth
	// logsdb *sql.DB

	// settings string
	settings *settings.Settings

	// subscriptionsBroker *subscriptions.Broker
	// logger *slog.Logger
}

// BaseAppConfig defines a BaseApp configuration option
type BaseAppConfig struct {
	IsDev   bool
	DataDir string
}

// NewBaseApp creates and returns a new BaseApp instance
// configured with the provided arguments.
//
// To initialize the app, you need to call `app.Bootstrap()`.
func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		isDev:   config.IsDev,
		dataDir: config.DataDir,
	}

	return app
}

// IsBootstrapped checks if the application was initialized
// (aka. whether Bootstrap() was called).
func (app *BaseApp) IsBootstrapped() bool {
	// return app.db != nil && app.logsdb != nil && app.settings != nil
	return app.db != nil
}

// Bootstrap initializes the application
// (aka. create data dir, open db connections, load settings, etc.).
//
// It will call ResetBootstrapState() if the application was already bootstrapped.
func (app *BaseApp) Bootstrap() error {
	// event := &BootstrapEvent{app}

	// if err := app.OnBeforeBootstrap().Trigger(event); err != nil {
	// 	return err
	// }

	// // clear resources of previous core state (if any)
	// if err := app.ResetBootstrapState(); err != nil {
	// 	return err
	// }

	// // ensure that data dir exist
	// if err := os.MkdirAll(app.DataDir(), os.ModePerm); err != nil {
	// 	return err
	// }

	if err := app.initAuthToken(); err != nil {
		return err
	}

	if err := app.initDataDB(); err != nil {
		return err
	}

	if err := app.initDataStore(); err != nil {
		return err
	}
	// if err := app.initLogsDB(); err != nil {
	// 	return err
	// }

	// if err := app.initLogger(); err != nil {
	// 	return err
	// }

	// // we don't check for an error because the db migrations may have not been executed yet
	// app.RefreshSettings()

	// // cleanup the pb_data temp directory (if any)
	// os.RemoveAll(filepath.Join(app.DataDir(), LocalTempDirName))

	// return app.OnAfterBootstrap().Trigger(event)

	return nil
}

// DB returns the default app database instance.
func (app *BaseApp) DB() *sql.DB {

	return app.db
}

func (app *BaseApp) DB_Context() context.Context {

	return app.db_context
}

// // LogsDB returns the app logs database instance.
// func (app *BaseApp) LogsDB() *sql.DB {
// 	return app.logsdb
// }

// DataDir returns the app data directory path.
func (app *BaseApp) DataDir() string {
	return app.dataDir
}

// IsDev returns whether the app is in dev mode.
//
// When enabled logs, executed sql statements, etc. are printed to the stderr.
func (app *BaseApp) IsDev() bool {
	return app.isDev
}

// Settings returns the loaded app settings.
func (app *BaseApp) Settings() *settings.Settings {
	return app.settings
}

// Store returns the app internal runtime store.
func (app *BaseApp) Store() *redka.DB {
	return app.store
}

// Store returns the app internal runtime store.
func (app *BaseApp) TokenAuth() *jwtauth.JWTAuth {
	return app.tokenAuth
}

func (app *BaseApp) initDataDB() error {

	ctx := context.Background()

	println("db start init")

	db_path := filepath.Join(app.DataDir(), "gk.db")
	db, err := sql.Open("sqlite3", "file:"+db_path+"?_journal=WAL")
	if err != nil {
		log.Fatal(err)
	}

	_ = models.New(db)

	// create tables
	// println(sqlz.Schema)
	if _, err := db.ExecContext(ctx, sqlz.Schema); err != nil {
		println("Creating Table Error :", err)
	}

	app.db = db
	app.db_context = ctx

	println("db end init")

	// TODO : ADD DEINIT TO CLOSE DB!!

	return nil
}

func (app *BaseApp) initDataStore() error {

	println("store start init")

	store_path := filepath.Join(app.DataDir(), "rk.db")
	store, err := redka.Open("file:"+store_path+"?_journal=WAL", nil)
	if err != nil {
		log.Fatal(err)
	}

	app.store = store

	println("store end init")

	// TODO : ADD DEINIT TO CLOSE STORE!!

	return nil
}

func (app *BaseApp) initAuthToken() error {

	println("tokenAuth start init")

	// more info on env :
	// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env_authToken := os.Getenv("AUTH_TOKEN")
	if env_authToken == "" {
		log.Fatal("Error: AUTH_TOKEN Not Found Inside .env file")
	}

	tokenAuth := jwtauth.New("HS256", []byte(env_authToken), nil, jwt.WithAcceptableSkew(30*time.Second))

	app.tokenAuth = tokenAuth

	println("tokenAuth end init")

	return nil
}
