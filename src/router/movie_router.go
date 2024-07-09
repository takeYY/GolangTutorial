package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"golang_tutorial/src/model"
	"golang_tutorial/src/repository"
)


type repoes struct {
	movieRepo repository.MovieRepository
}


func GetMovie(c echo.Context) error {
	log.Info("accessed get movie")

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

	repoes := repoes{
		movieRepo: repository.MovieRepository{Session: tx},
	}

	id := 1
	movie, err := repoes.movieRepo.GetMovieById(&id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "movie not found"})
	}

	log.Info(fmt.Printf("first movie is %s", movie.Title))

	return c.JSON(http.StatusOK, movie)
}
