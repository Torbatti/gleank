package views

import (
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

type UserInfo struct {
	isLoggedIn bool
	name       string
}

type PageInfo struct {
	// seo related
	HeadInfo HeadInfo

	UserInfo UserInfo
}

func BindViews(app core.App, r *chi.Mux) {
	hxIslandUserHeader(app, r)

	BindStaticPublic(app, r)

	indexPage(app, r)

	authNewPage(app, r)
	authPages(app, r)

	folderPage(app, r)

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
