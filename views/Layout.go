package views

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/go-chi/chi/v5"
)

func hxIslandUserHeader(app core.App, r *chi.Mux) {

	r.Get("/island/user-header", func(w http.ResponseWriter, r *http.Request) {

		var user_name string

		var is_not_logged_in bool
		cookie, err := r.Cookie("jwt")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				is_not_logged_in = true
				// http.Error(w, "cookie not found", http.StatusBadRequest)
			default:
				fmt.Println("server error :", err)
				// http.Error(w, "server error", http.StatusInternalServerError)
			}
		}
		// println(errors.Is(err, http.ErrNoCookie))
		println("is logged in : ", !is_not_logged_in)
		if !is_not_logged_in {

			jwtToken := app.TokenAuth()
			token, err := jwtToken.Decode(cookie.Value)
			if err != nil {
				fmt.Println("Decoding Error:", err)
			}

			claim, ok := token.Get("user_name")
			if !ok {
				println("claim is empty")
			}
			fmt.Println(fmt.Sprint(claim))

			user_name = fmt.Sprint(claim)
		}
		if is_not_logged_in {
			user_name = ""
		}

		info := PageInfo{
			UserInfo: UserInfo{
				isLoggedIn: !is_not_logged_in,
				name:       user_name,
			},
		}

		var buffer = bytes.NewBufferString("")
		IslandHeader(info.UserInfo).Render(r.Context(), buffer)
		// app.Store().Str().Set("page-index", buffer.String())

		w.Write([]byte(buffer.String()))
	})

}
