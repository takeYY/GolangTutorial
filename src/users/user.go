package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)


func GetUserName(c echo.Context) error {
	log.Info("accessed get user name")

	name := c.Param("name")
	r, err := GetUser(&name)
	if err != nil {
		log.Error(err)
	}

	log.Info(fmt.Printf("return response %s", r))
	return c.JSON(http.StatusOK, r)
}
