package repository

import (
	"golang_tutorial/src/model"

	"gorm.io/gorm"
)


type MovieRepository struct {
	Session *gorm.DB
}


func (r *MovieRepository) GetMovieById(id *int) (*model.Movie, error) {
	var result model.Movie

	if err := r.Session.Where("id IN (?)", *id).First(&result).Error; err  != nil {
		return nil, err
	}

	return &result, nil
}
