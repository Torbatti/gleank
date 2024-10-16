package views

import (
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/go-chi/chi/v5"
)

// https://www.alexedwards.net/blog/serving-static-sites-with-go
func BindStaticPublic(app core.App, r *chi.Mux) {
	// r.Group(func(r chi.Router) {
	// 	r.Use(middleware.RequestID)
	// 	r.Use(middleware.Logger)
	// 	r.Use(middleware.Recoverer)
	// })

	static := http.FileServer(http.Dir("./views/public"))

	// r.Mount("/", static)
	r.Mount("/public/", http.StripPrefix("/public/", static))
}
