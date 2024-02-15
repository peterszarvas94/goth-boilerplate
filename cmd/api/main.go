package main

import (
	"fmt"
	"goth/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
