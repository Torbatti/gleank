package views

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/Torbatti/gleank/core"
	"github.com/Torbatti/gleank/models"
	"github.com/go-chi/chi/v5"
)

func authPages(app core.App, r *chi.Mux) {

	r.Get("/auth/register", func(w http.ResponseWriter, r *http.Request) {

		info := PageInfo{
			HeadInfo: HeadInfo{
				Title:       "Register",
				Description: "Register Page",
			},
		}

		var buffer = bytes.NewBufferString("")
		auth_register_page(info).Render(r.Context(), buffer)
		app.Store().Str().Set("page-auth_register", buffer.String())

		w.Write([]byte(buffer.String()))
	})

	r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {

		info := PageInfo{
			HeadInfo: HeadInfo{
				Title:       "Login",
				Description: "Login Page",
			},
		}

		var buffer = bytes.NewBufferString("")
		auth_login_page(info).Render(r.Context(), buffer)
		app.Store().Str().Set("page-auth_login", buffer.String())

		w.Write([]byte(buffer.String()))
	})

	r.Post("/hx/auth/logout", func(w http.ResponseWriter, r *http.Request) {

		// Delete jwt cookie
		jwt_cookie := http.Cookie{
			Name:     "jwt",
			Path:     "/",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &jwt_cookie)

		w.Header().Set("HX-Redirect", "/")
	})

	r.Post("/hx/auth/register", func(w http.ResponseWriter, r *http.Request) {

		user_name := r.PostFormValue("user_name")
		email := r.PostFormValue("email")

		// Check wheter email or username is empty
		if user_name == "" || email == "" {
			w.Write([]byte("Email or Username is either empty"))
		}

		ctx := context.Background()
		queries := models.New(app.DB())

		// check wheter username is already exist
		fetched_username, err := queries.GetUserByName(ctx, user_name)
		if err != nil {
			fmt.Println(err)
		}
		if fetched_username != (models.User{}) {
			hx, err := app.Store().Str().Get("hx-register_error_box")
			if err != nil {
				fmt.Println(err)
			}
			if hx.String() == "" {
				var buffer = bytes.NewBufferString("")
				register_errorr_box().Render(r.Context(), buffer)
				app.Store().Str().Set("hx-register_error_box", buffer.String())
				w.Write([]byte(buffer.String()))
			}
			w.Write([]byte(hx.String()))
		}

		// check wheter email is already exist
		fetched_email, err := queries.GetUser(ctx, 12)
		if err != nil {
			fmt.Println(err)
		}
		if fetched_email != (models.User{}) {
			hx, err := app.Store().Str().Get("hx-register_error_box")
			if err != nil {
				fmt.Println(err)
			}
			if hx.String() == "" {
				var buffer = bytes.NewBufferString("")
				register_errorr_box().Render(r.Context(), buffer)
				app.Store().Str().Set("hx-register_error_box", buffer.String())
				w.Write([]byte(buffer.String()))
			}
			w.Write([]byte(hx.String()))
		}

		jwtToken := app.TokenAuth()
		_, tokenString, _ := jwtToken.Encode(map[string]interface{}{"user_name": user_name})

		jwt_cookie := http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &jwt_cookie)

		// session_cookie := http.Cookie{
		// 	Name:     "session",
		// 	Value:    "",
		// 	Path:     "/",
		// 	MaxAge:   3600,
		// 	HttpOnly: true,
		// 	Secure:   true,
		// 	SameSite: http.SameSiteLaxMode,
		// }

		// queries.

		// http.SetCookie(w, &session_cookie)
		w.Header().Set("HX-Redirect", "/")
	})
}
