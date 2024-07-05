package main

import (
	"golang_tutorial/src/apps/hello"
	"golang_tutorial/src/apps/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.Hello)
	e.GET("/users/:name", users.GetUserName)

	e.Logger.Fatal(e.Start(":8000"))
}
