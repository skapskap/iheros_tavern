package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	registerRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.port)))
}