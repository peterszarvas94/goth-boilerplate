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

	msg := fmt.Sprintf("server started on port %s", server.Addr)
	fmt.Println(msg)

	logger.Info(msg)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
