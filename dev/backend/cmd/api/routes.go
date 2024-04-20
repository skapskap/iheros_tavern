package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skapskap/iheros_tavern/internal/data"
	"github.com/skapskap/iheros_tavern/internal/handlers"
)

func registerRoutes(e *echo.Echo, db data.DBTX) {
	v1 := e.Group("/v1")

    v1.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "running"})
    })

	v1.POST("/auth/register", func(c echo.Context) error {
		return handlers.AuthRegister(c, db)
	})

	v1.POST("/auth/login", func(c echo.Context) error {
		return handlers.AuthLogin(c, db)
	})
}