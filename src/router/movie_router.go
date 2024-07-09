package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"golang_tutorial/src/repository"
)



type Movie struct {
	ID 				int 						`gorm:"primary_key;column:id"`
	CreatedAt time.Time 			`gorm:"column:created_at"`
	UpdatedAt time.Time 			`gorm:"column:updated_at"`
	Title 		string 					`gorm:"column:title"`
	Overview 	sql.NullString 	`gorm:"column:overview"`
}

func GetMovie(c echo.Context) error {
	log.Info("accessed get movie")

	dbSess := repository.ConnectDB()
	tx := dbSess.Begin()

	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var mov Movie
	tx.First(&mov)

	log.Info(fmt.Printf("first movie is %s", mov.Title))

	return c.JSON(http.StatusOK, mov)
}
