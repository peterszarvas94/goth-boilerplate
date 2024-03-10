package handlers

// no found handler

import (
	"goth/web/templates/pages"
	"net/http"

	"github.com/a-h/templ"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	component := pages.NotFound()
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}
