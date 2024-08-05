package repository

import (
	"golang_tutorial/src/model"
	"golang_tutorial/src/query"

	"gorm.io/gorm"
)

type MovieRepository struct {
	Session *gorm.DB
}

func (r *MovieRepository) GetMovieById(id *int32) (*model.Movie, error) {
	m := query.Use(r.Session).Movie

	movie, e := m.Preload(m.Genres).Where(m.ID.Eq(*id)).First()
	if e != nil {
		return nil, e
	}

	return movie, nil
}

func (r *MovieRepository) GetMovies() ([]*model.Movie, error) {
	m := query.Use(r.Session).Movie

	movie, err := m.Preload(m.Genres).Find()
	if err != nil {
		return nil, err
	}

	return movie, nil
}
