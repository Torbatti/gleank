package apis

import (
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/Torbatti/gleank/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitApi(app core.App) (*chi.Mux, error) {
	r := chi.NewRouter()

	// Default Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	// Bindings
	views.BindViews(app, r)

	return r, nil
}
