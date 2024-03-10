package server

import (
	"goth/internal/server/handlers"
	"goth/internal/server/middlewares"
	"net/http"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	/* fallback */
	mux.HandleFunc("/", handlers.NotFound)

	/* static files */
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	/* routes */
	mux.HandleFunc("GET /{$}", handlers.Index)
	mux.HandleFunc("GET /users/{$}", middlewares.DB(handlers.GetUsers))
	mux.HandleFunc("GET /signup/{$}", handlers.SignupPage)
	mux.HandleFunc("POST /signup/{$}", middlewares.DB(handlers.Signup))
	mux.HandleFunc("GET /signin/{$}", handlers.SigninPage)
	mux.HandleFunc("POST /signin/{$}", middlewares.DB(handlers.Signin))
	mux.HandleFunc("GET /dashboard/{$}", handlers.Dashboard)

	return mux
}
