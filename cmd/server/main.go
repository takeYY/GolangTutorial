package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang_tutorial/config"
	"golang_tutorial/internal/movie"
)

func main() {
	// 設定ファイルの読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ハンドラの設定
	movieHandler := movie.NewHandler(&cfg.ReaderDatabase, &cfg.WriterDatabase)
	movieHandler.RegisterRoutes(e)

	// サーバーの起動
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Starting server on %s", serverAddr)
	e.Logger.Fatal(e.Start(serverAddr))
}
