package handlers

import (
	"goth/web/templates/pages"
	"net/http"

	"github.com/a-h/templ"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	component := pages.Dashboard();
	handler := templ.Handler(component);
	handler.ServeHTTP(w, r);
}
