package server

import (
	h "goth/internal/server/handlers"
	m "goth/internal/server/middlewares"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.NotFoundHandler)
	mux.HandleFunc("GET /{$}", h.HelloWorldHandler)
	mux.HandleFunc("GET /users", m.DBMiddleware(h.GetUsersHandler))

	return mux
}
