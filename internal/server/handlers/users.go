package handlers

import (
	"encoding/json"
	"fmt"
	"goth/internal/slogger"
	"goth/internal/database"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(database.Service)
	if !ok {
		slogger.Log.Fatal("error getting db from context")
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(db.GetUsers())

	if err != nil {
		slogger.Log.Fatal(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
	}

	_, _ = w.Write(jsonResp)
}
