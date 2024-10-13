package views

import (
	"bytes"
	"context"
	"log"
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/Torbatti/gleank/models"
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
	indexPage(app, r)

	authNewPage(app, r)

	folderPage(app, r)
}

func indexPage(app core.App, r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		// // Checks if page is cached in app.store
		// page, err := app.Store().Str().Get("page-index")
		// if err != nil {
		// 	log.Println(err)
		// }
		// if page.String() != "" {
		// 	w.Write([]byte(page))
		// }

		info := PageInfo{
			HeadInfo: HeadInfo{
				Title:       "Gleank",
				Description: "a Link Shortener Social Media",
			},
		}

		var buffer = bytes.NewBufferString("")
		index_page(info).Render(r.Context(), buffer)
		app.Store().Str().Set("page-index", buffer.String())

		w.Write([]byte(buffer.String()))
	})
}

func authNewPage(app core.App, r *chi.Mux) {
	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		q_name := r.URL.Query().Get("name")

		if q_name == "" {
			w.Write([]byte("name not been set"))
		} else {
			w.Write([]byte(q_name))
		}

	})
}

func folderPage(app core.App, r *chi.Mux) {
	r.Get("/folder/{path}", func(w http.ResponseWriter, r *http.Request) {
		pathParam := chi.URLParam(r, "date")

		// Checks if page is cached in app.store
		page, err := app.Store().Str().Get("folder-" + pathParam)
		if err != nil {
			log.Println(err)
		}
		if len(page) > 0 {
			w.Write([]byte(page))
		}

		// Retrive folder data
		ctx := context.Background()
		queries := models.New(app.DB())

		// get the author we just inserted
		fetchedAuthor, err := queries.GetFolderByPath(ctx, pathParam)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte(fetchedAuthor.Name))

	})
}
