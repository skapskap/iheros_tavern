package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

    v1.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "running"})
    })
}