package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/uocli/go-microservice/internal/database"
	"log"
	"net/http"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	err := server.registerRoutes()
	if err != nil {
		return nil
	}
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() error {
	return nil
}
