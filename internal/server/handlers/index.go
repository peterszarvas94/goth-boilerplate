package handlers

import (
	"goth/web/templates/pages"
	"net/http"

	"github.com/a-h/templ"
)

func Index(w http.ResponseWriter, r *http.Request) {
	component := pages.Index();
	handler := templ.Handler(component);
	handler.ServeHTTP(w, r);
}
