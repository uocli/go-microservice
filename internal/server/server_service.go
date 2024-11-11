package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	customers, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}
