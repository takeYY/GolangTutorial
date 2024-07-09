package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang_tutorial/src/hello"
	"golang_tutorial/src/movie"
	"golang_tutorial/src/router"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.Hello)
	e.GET("/users/:name", router.GetUserName)
	e.GET("movies/get", movie.GetMovie)

	e.Logger.Fatal(e.Start(":8000"))
}
