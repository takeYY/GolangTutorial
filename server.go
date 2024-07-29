package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang_tutorial/src/handlers"
	"golang_tutorial/src/hello"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.Hello)
	e.GET("/users/:name", handlers.GetUserName)
	e.GET("movies/get/:id", handlers.GetMovie)

	e.Logger.Fatal(e.Start(":8000"))
}
