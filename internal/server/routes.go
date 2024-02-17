package server

import (
	"encoding/json"
	"fmt"
	"goth/internal/slogger"
	"goth/web/templates/pages"
	"net/http"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/users", s.getUsersHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	component := pages.Index();
	handler := templ.Handler(component);
	handler.ServeHTTP(w, r);
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		slogger.Log.Fatal(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.GetUsers())

	if err != nil {
		slogger.Log.Fatal(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
	}

	_, _ = w.Write(jsonResp)
}
