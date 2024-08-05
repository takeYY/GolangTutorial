package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

	i64, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	id := int32(i64) // ParseInt が int64になるので再度キャストする

	movieWithGenres, err := repoes.movieRepo.GetMovieById(&id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "movie not found"})
	}

	log.Info(fmt.Printf("first movie is %s", movieWithGenres.Title))

	return c.JSON(http.StatusOK, movieWithGenres)
}

func GetMovies(c echo.Context) error {
	log.Info("accessed get movies")

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

	moviesWithGenres, err := repoes.movieRepo.GetMovies()
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "movies are not found"})
	}

	return c.JSON(http.StatusOK, moviesWithGenres)
}
