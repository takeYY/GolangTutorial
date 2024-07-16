package repository

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	default_mysql "github.com/go-sql-driver/mysql"
)


func ConnectDB() *gorm.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Error(fmt.Printf("failed to load location %s", err))
	}

	c := default_mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PWD"),
		Addr:      os.Getenv("DB_HOST"),
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
