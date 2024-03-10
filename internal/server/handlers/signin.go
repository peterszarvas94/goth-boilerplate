package handlers

import (
	"fmt"
	"goth/internal/database"
	"goth/internal/database/schema"
	"goth/web/templates/pages"
	"html"
	"net/http"

	"github.com/a-h/templ"
)

func SigninPage(w http.ResponseWriter, r *http.Request) {
	component := pages.Signin()
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(database.Service)
	if !ok {
		logger.Error("error getting db from context")
		InternalServerError(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		logger.Error("error parsing form")
		InternalServerError(w, r)
		return
	}

	form := r.Form

	usernameOrEmail := html.EscapeString(form.Get("username_or_email"))
	if usernameOrEmail == "" {
		logger.Error("username_or_email is empty")
		BadRequest(w, r)
		return
	}

	password := html.EscapeString(form.Get("password"))
	if password == "" {
		logger.Error("password is empty")
		BadRequest(w, r)
		return
	}

	userPropsSignin := schema.UserPropsSignin{
		UsernameOrEmail: usernameOrEmail,
		Password:        password,
	}

	user, err := db.UserSignin(userPropsSignin)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	// TODO: set session and remove next line
	logger.Info(fmt.Sprintf("user signed in: %s", user.Id))

	http.Redirect(
		w,
		r,
		"/dashboard",
		http.StatusSeeOther,
	)
}
