package main

import (
	"golang_tutorial/src/repository"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./src/query",
		Mode:    gen.WithoutContext,
	})

	gormdb := repository.ConnectDB()

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	all := g.GenerateAllTable() // database to table model

	// create movie with genres
	genres_neo := g.GenerateModel("genre_neo")
	movie_neo := g.GenerateModel("movie_neo", gen.FieldRelate(field.HasMany, "Genres_neo", genres_neo, &field.RelateConfig{
		GORMTag: field.GormTag{"foreignKey": []string{"MovieRefer"}, "references": []string{"ID"}},
	}))

	// TODO: movie と genre の多対多の関係を作りたい...
	// TODO: 作ったができているかわからんから確認すること!!
	genres := g.GenerateModel("genre")
	movie := g.GenerateModel("movie", gen.FieldRelate(field.Many2Many, "Genres", genres, &field.RelateConfig{
		RelateSlice: true,
		GORMTag:     field.GormTag{"many2many": []string{"movie_genre"}},
	}))

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(genres_neo, movie_neo, movie)
	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
