package apis

import (
	"net/http"
	"time"

	"github.com/Torbatti/gleank/core"
)

// ServeConfig defines a configuration struct for apis.Serve().
type ServeConfig struct {
	// ShowStartBanner indicates whether to show or hide the server start console message.
	ShowStartBanner bool

	// HttpAddr is the TCP address to listen for the HTTP server (eg. `127.0.0.1:80`).
	HttpAddr string

	// HttpsAddr is the TCP address to listen for the HTTPS server (eg. `127.0.0.1:443`).
	HttpsAddr string

	// Optional domains list to use when issuing the TLS certificate.
	//
	// If not set, the host from the bound server address will be used.
	//
	// For convenience, for each "non-www" domain a "www" entry and
	// redirect will be automatically added.
	CertificateDomains []string

	// AllowedOrigins is an optional list of CORS origins (default to "*").
	AllowedOrigins []string
}

func Serve(app core.App, config ServeConfig) (*http.Server, error) {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	// // ensure that the latest migrations are applied before starting the server
	// if err := runMigrations(app); err != nil {
	// 	return nil, err
	// }

	// // reload app settings in case a new default value was set with a migration
	// // (or if this is the first time the init migration was executed)
	// if err := app.RefreshSettings(); err != nil {
	// 	color.Yellow("=====================================")
	// 	color.Yellow("WARNING: Settings load error! \n%v", err)
	// 	color.Yellow("Fallback to the application defaults.")
	// 	color.Yellow("=====================================")
	// }

	router, err := InitApi(app)
	if err != nil {
		return nil, err
	}

	// // configure cors
	// router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	AllowOrigins: config.AllowedOrigins,
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// }))

	// start http server
	// ---
	mainAddr := config.HttpAddr
	if config.HttpsAddr != "" {
		mainAddr = config.HttpsAddr
	}

	server := &http.Server{
		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
		// WriteTimeout: 60 * time.Second, // breaks sse!
		Handler: router,
		Addr:    mainAddr,
	}

	// // start HTTPS server
	// if config.HttpsAddr != "" {
	// 	// if httpAddr is set, start an HTTP server to redirect the traffic to the HTTPS version
	// 	if config.HttpAddr != "" {
	// 		go http.ListenAndServe(config.HttpAddr, certManager.HTTPHandler(nil))
	// 	}

	// 	return server, server.ListenAndServeTLS("", "")
	// }

	// OR start HTTP server
	return server, server.ListenAndServe()
}
