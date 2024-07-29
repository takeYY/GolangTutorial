package repository

import (
	"golang_tutorial/src/model"
	"golang_tutorial/src/query"

	"gorm.io/gorm"
)

type GenreRepository struct {
	Session *gorm.DB
}

func (r *GenreRepository) GetGenreByCode(code *string) (*model.Genre, error) {
	g := query.Use(r.Session).Genre

	result, err := g.Where(g.Code.Eq(*code)).First()
	if err != nil {
		return nil, err
	}

	return result, nil
}
