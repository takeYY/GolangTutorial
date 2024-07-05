package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// User
type User struct {
  Name  string `json:"name" xml:"name"`
  Description string `json:"description" xml:"description"`
}

func GetUserName(c echo.Context) error {
	log.Info("accessed get user name")

	name := &User{
		Name: c.Param("name"),
		Description: "説明のテストですよ",
	}

	log.Info(fmt.Printf("return response %s", name))
	return c.JSON(http.StatusOK, name)
}
