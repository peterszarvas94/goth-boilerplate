package middlewares

import (
	"context"
	"goth/internal/database"
	"net/http"
)

func DBMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", database.DB)
		r = r.WithContext(ctx)
		next(w, r)
	})
}
