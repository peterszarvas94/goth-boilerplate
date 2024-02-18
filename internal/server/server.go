package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goth/internal/config"
	"goth/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(config.App.PORT)
	if err != nil {
		return nil, err
	}

	// seed the database
	database.DB.Seed()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
