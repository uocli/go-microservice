package server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/uocli/go-microservice/internal/database"
	"github.com/uocli/go-microservice/internal/models"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Readiness(echo.Context) error
	Liveness(echo.Context) error
	GetAllCustomers(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	GetAllServices(ctx echo.Context) error
	GetAllVendors(ctx echo.Context) error
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
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)

	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)

	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)

	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)

	return nil
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	if !s.DB.Ready() {
		return ctx.JSON(http.StatusInternalServerError, models.Health{
			Status: "Failure",
		})
	}
	return ctx.JSON(http.StatusOK, models.Health{
		Status: "OK",
	})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{
		Status: "OK",
	})
}
