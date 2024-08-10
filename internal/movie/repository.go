package movie

import (
	"golang_tutorial/model"
	"golang_tutorial/query"

	"gorm.io/gorm"
)

// 映画データのリポジトリインターフェース
type (
	IQueryRepository interface {
		FindByID(id *int32) (*model.Movie, error)
		Find() ([]*model.Movie, error)
	}
	ICommandRepository interface {
		Save(target *NewMovie) (*model.Movie, error)
		Delete(id *string) error
	}
	MovieRepository struct {
		Session *gorm.DB
	}
)

func (r *MovieRepository) FindByID(id *int32) (*model.Movie, error) {
	m := query.Use(r.Session).Movie

	movie, e := m.Preload(m.Genres).Where(m.ID.Eq(*id)).First()
	if e != nil {
		return nil, e
	}

	return movie, nil
}

func (r *MovieRepository) Find() ([]*model.Movie, error) {
	m := query.Use(r.Session).Movie

	movies, e := m.Preload(m.Genres).Find()
	if e != nil {
		return nil, e
	}

	return movies, nil
}

func (r *MovieRepository) Save(target *NewMovie) (*model.Movie, error) {
	var err error
	defer func() {
		if err != nil {
			r.Session.Rollback()
		} else {
			r.Session.Commit()
		}
	}()

	m := query.Use(r.Session).Movie

	movie := model.Movie{
		Title:    target.Title,
		Overview: target.Overview,
		Genres:   []model.Genre{},
	}
	err = m.Create(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) Delete(id *string) error {
	return nil
}
