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
	}
	ICommandRepository interface {
		Save(movie *model.Movie) error
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

func (r *MovieRepository) Save(movie *model.Movie) error {
	return nil
}

func (r *MovieRepository) Delete(id *string) error {
	return nil
}
