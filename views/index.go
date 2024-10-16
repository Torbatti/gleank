package views

import (
	"bytes"
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/go-chi/chi/v5"
)

func indexPage(app core.App, r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		// // Checks if page is cached in app.store
		// page, err := app.Store().Str().Get("page-index")
		// if err != nil {
		// 	log.Println(err)
		// }
		// // save it to not call the function twice
		// page_string := page.String()
		// // if in store -> Great
		// if page_string != "" {
		// 	w.Write([]byte(page))
		// }

		// // if not in store -> make the page
		// if page_string == "" {

		// 	info := PageInfo{
		// 		HeadInfo: HeadInfo{
		// 			Title:       "Gleank",
		// 			Description: "a Link Shortener Social Media",
		// 		},
		// 	}

		// 	var buffer = bytes.NewBufferString("")
		// 	index_page(info).Render(r.Context(), buffer)
		// 	app.Store().Str().Set("page-index", buffer.String())

		// 	w.Write([]byte(buffer.String()))

		// }

		info := PageInfo{
			HeadInfo: HeadInfo{
				Title:       "Gleank",
				Description: "a Link Shortener Social Media",
			},
			UserInfo: UserInfo{
				isLoggedIn: false,
				name:       "Torbatti",
			},
		}

		var buffer = bytes.NewBufferString("")
		index_page(info).Render(r.Context(), buffer)
		app.Store().Str().Set("page-index", buffer.String())

		w.Write([]byte(buffer.String()))
	})
}
