package main

import (
	"fmt"
	"goth/internal/server"
	"goth/internal/slogger"
)

func main() {
	logger := slogger.Get()

	server, err := server.NewServer()
	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	logger.Info("Server started")

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
