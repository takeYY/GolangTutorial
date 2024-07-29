package repository

import (
	"golang_tutorial/src/model"
	"golang_tutorial/src/query"

	"gorm.io/gorm"
)

type MovieRepository struct {
	Session *gorm.DB
}

func (r *MovieRepository) GetMovieById(id *int32) (*model.MovieWithGenres, error) {
	m := query.Use(r.Session).Movie
	g := query.Use(r.Session).Genre
	mg := query.Use(r.Session).MovieGenre

	type Result struct {
		Title    string
		Overview string
		Name     string
	}
	var results []Result

	err := mg.Where(mg.MovieID.Eq(*id)).Select(m.Title, m.Overview, g.Name).LeftJoin(m, m.ID.EqCol(mg.MovieID)).LeftJoin(g, g.ID.EqCol(mg.GenreID)).Scan(&results)
	if err != nil {
		return nil, err
	}

	var genres []string
	for _, result := range results {
		genres = append(genres, result.Name)
	}

	movieWithGenre := model.MovieWithGenres{
		Title:    results[0].Title,
		Overview: results[0].Overview,
		Genres:   genres,
	}

	return &movieWithGenre, nil
}

func (r *MovieRepository) GetMovies() ([]*model.Movie, error) {
	m := query.Use(r.Session).Movie

	result, err := m.Find()
	if err != nil {
		return nil, err
	}

	return result, nil
}
