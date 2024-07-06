package movie

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	default_mysql "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)


func ConnectDB() *gorm.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Error(fmt.Printf("failed to load location %s", err))
	}

	c := default_mysql.Config{
		DBName:    "db4test",
		User:      "user",
		Passwd:    "password",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		AllowNativePasswords: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Error(fmt.Printf("failed to connect mysql %s", err))
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Error(fmt.Printf("failed to use gorm %s", err))
	}

	return gormDB
}



type Movie struct {
	ID 				int 						`gorm:"primary_key;column:id"`
	CreatedAt time.Time 			`gorm:"column:created_at"`
	UpdatedAt time.Time 			`gorm:"column:updated_at"`
	Title 		string 					`gorm:"column:title"`
	Overview 	sql.NullString 	`gorm:"column:overview"`
}

func GetMovie(c echo.Context) error {
	log.Info("accessed get movie")

	dbSess := ConnectDB()
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
