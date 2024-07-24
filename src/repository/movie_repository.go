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

	result, err := m.Where(m.ID.Eq(*id)).First()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MovieRepository) GetMovies() ([]*model.Movie, error) {
	m := query.Use(r.Session).Movie

	result, err := m.Find()
	if err != nil {
		return nil, err
	}

	return result, nil
}
