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

type Server struct {
	port int
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(config.App.PORT)
	if err != nil {
		return nil, err
	}

	NewServer := &Server{
		port: port,
	}

	// seed the database
	database.DB.Seed()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
