package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:name", getUserName)
	e.Logger.Fatal(e.Start(":8000"))
}

// User
type User struct {
  Name  string `json:"name" xml:"name"`
  Description string `json:"description" xml:"description"`
}

func getUserName(c echo.Context) error {
	name := &User{
		Name: c.Param("name"),
		Description: "説明のテストですよ",
	}

	return c.JSON(http.StatusOK, name)
}
