package main

import (
	"golang_tutorial/src/repository"

	"gorm.io/gen"
)

// TODO: movie と genre を結合した構造体を自動生成酢量にしたい
// refs: https://zenn.dev/rescuenow/articles/da1cb5f574fb0c
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../src/query",
		Mode:    gen.WithoutContext,
	})

	gormdb := repository.ConnectDB()

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	all := g.GenerateAllTable() // database to table model

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
