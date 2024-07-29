package handlers

import (
	"golang_tutorial/src/model"
	"golang_tutorial/src/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetGenreByCode(c echo.Context) error {
	log.Info("accessed get genre by code")

	dbSession := repository.ConnectDB()
	tx := dbSession.Begin()

	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	repoes := repository.GenreRepository{Session: tx}

	code := c.Param("code")

	genre, err := repoes.GetGenreByCode(&code)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "genre not found"})
	}

	return c.JSON(http.StatusOK, genre)
}
