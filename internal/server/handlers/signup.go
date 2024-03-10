package handlers

import (
	"fmt"
	"goth/internal/database"
	"goth/internal/database/schema"
	"goth/internal/uuid"
	"goth/web/templates/pages"
	"html"
	"net/http"

	"github.com/a-h/templ"
)

func SignupPage(w http.ResponseWriter, r *http.Request) {
	component := pages.Signup()
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(database.Service)
	if !ok {
		InternalServerError(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		InternalServerError(w, r)
		return
	}

	form := r.Form

	username := html.EscapeString(form.Get("username"))
	if username == "" {
		BadRequest(w, r)
		return
	}

	email := html.EscapeString(form.Get("email"))
	if email == "" {
		BadRequest(w, r)
		return
	}

	password := html.EscapeString(form.Get("password"))
	if password == "" {
		BadRequest(w, r)
		return
	}

	id, err := uuid.New("usr")
	if err != nil {
		InternalServerError(w, r)
		return
	}

	userPropsSignup := schema.UserPropsSignup{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	err = db.UserSignup(userPropsSignup)
	if err != nil {
		InternalServerError(w, r)
		return
	}

	logger.Info(fmt.Sprintf("user signed up as %s (%s)", username, email))

	http.Redirect(
		w,
		r,
		"/signin",
		http.StatusSeeOther,
	)
}
