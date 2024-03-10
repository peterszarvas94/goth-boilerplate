package handlers

import (
	"encoding/json"
	"fmt"
	"goth/internal/slogger"
	"goth/internal/database"
	"net/http"
)

var logger = slogger.Get()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(database.Service)
	if !ok {
		logger.Error("error getting db from context")
		InternalServerError(w, r)
		return
	}

	users, err := db.UserGetAll()
	if err != nil {
		logger.Error(fmt.Sprintf("error getting users from db. Err: %v", err))
		InternalServerError(w, r)
		return
	}

	jsonResp, err := json.Marshal(users)
	if err != nil {
		logger.Error(fmt.Sprintf("error handling JSON marshal. Err: %v", err))
		InternalServerError(w, r)
		return
	}

	w.Write(jsonResp)
}
