package handlers

import (
	"encoding/json"
	"fmt"
	"goth/internal/slogger"
	"goth/internal/database"
	"net/http"
)

var logger = slogger.Get()

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(database.Service)
	if !ok {
		logger.Error("error getting db from context")
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(db.GetUsers())

	if err != nil {
		logger.Error(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
	}

	w.Write(jsonResp)
}
