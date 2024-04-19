package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skapskap/iheros_tavern/internal/database"
)

type config struct {
	port int
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4869, "API server port")
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := database.SetupDatabase()
	if err != nil {
		e.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	registerRoutes(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.port)))
}