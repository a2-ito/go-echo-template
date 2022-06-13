package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"go-echo-template/config"
	"go-echo-template/helloworld"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, helloworld.Helloworld())
}

func main() {
	// Echo instance
	e := echo.New()

	// logger zap
	logger, _ := zap.NewDevelopment()

	// global
	zap.ReplaceGlobals(logger)

	logger.Info("Hello zap")

	goenv := config.LoadConfig()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	// Start server
	logger.Info("http server started on " + goenv.Port)
	e.Logger.Fatal(e.Start(":" + goenv.Port))
}
