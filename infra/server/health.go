package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

type HealthApi struct {
}

func NewHealth() *HealthApi {
	return &HealthApi{}
}

func (controller HealthApi) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.GET("/health", controller.Health)
}

func (controller HealthApi) Health(c echo.Context) error {
	response := map[string]string{
		"status": "UP",
	}
	return c.JSON(http.StatusOK, response)
}
