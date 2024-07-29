package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"golang_tutorial/src/model"
	"golang_tutorial/src/repository"
)

func GetUserName(c echo.Context) error {
	log.Info("accessed get user name")

	name := c.Param("name")
	r, err := repository.GetUser(&name)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			Message: fmt.Sprintf("%s is not found.", name),
		})
	}

	log.Info(fmt.Printf("return response %s", r))
	return c.JSON(http.StatusOK, r)
}
