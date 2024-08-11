package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"

	"golang_tutorial/config"
	"golang_tutorial/internal/movie"
)

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// 設定ファイルの読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	e := echo.New()
	e.Validator = NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ハンドラの設定
	movieHandler := movie.NewHandler(cfg)
	movieHandler.RegisterRoutes(e)

	// サーバーの起動
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Starting server on %s", serverAddr)
	e.Logger.Fatal(e.Start(serverAddr))
}
