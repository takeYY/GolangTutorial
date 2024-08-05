package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang_tutorial/src/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Hello)
	e.GET("/users/:name", handlers.GetUserName)
	e.GET("/movies/get/:id", handlers.GetMovie)
	e.GET("/movies", handlers.GetMovies)
	e.GET("/genres/:code", handlers.GetGenreByCode)

	e.Logger.Fatal(e.Start(":8000"))
}
