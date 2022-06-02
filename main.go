package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo-template/config"
	"go-echo-template/helloworld"
	"net/http"
)

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, helloworld.Helloworld())
}

func main() {

	// Echo instance
	e := echo.New()

	goenv := config.LoadConfig()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":" + goenv.Port))
}
