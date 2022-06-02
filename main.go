package main

import (
  "net/http"
  "go-echo-template/helloworld"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, helloworld.Helloworld())
}

func main()  {

    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", hello)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
