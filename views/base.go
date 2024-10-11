package views

import (
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/go-chi/chi/v5"
)

type HeadInfo struct {
	Title       string
	Description string
}

type PageInfo struct {
	// seo related
	HeadInfo HeadInfo

	// Game models.Game
}

func BindViews(app core.App, r *chi.Mux) {
	r.Get("/", indexPage(app))
}

func indexPage(w http.ResponseWriter, r *http.Request, app core.App) {

}
