package main

import (
	"github.com/uocli/go-microservice/internal/database"
	"github.com/uocli/go-microservice/internal/server"
	"log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to create database client: %s", err)
	}

	svr := server.NewEchoServer(db)
	if err = svr.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
