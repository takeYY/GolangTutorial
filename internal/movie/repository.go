package movie

import (
	"context"
	"golang_tutorial/config"
	"golang_tutorial/db"
	"golang_tutorial/model"
	"golang_tutorial/query"

	"gorm.io/gorm"
)

// 映画データのリポジトリインターフェース
type (
	IQueryRepository interface {
		FindByID(ctx context.Context, id *int32) (*model.Movie, error)
		Find(ctx context.Context) ([]*model.Movie, error)
	}
	ICommandRepository interface {
		Save(ctx context.Context, target *NewMovie) (*model.Movie, error)
		Delete(ctx context.Context, id *string) error
	}
	MovieRepository struct {
		dbCfg *config.Database
		db    *gorm.DB
	}
)

func (r *MovieRepository) FindByID(ctx context.Context, id *int32) (*model.Movie, error) {
	r.db = db.ConnectDB(r.dbCfg)
	m := query.Use(r.db).Movie

	movie, e := m.WithContext(ctx).Preload(m.Genres).Where(m.ID.Eq(*id)).First()
	if e != nil {
		return nil, e
	}

	return movie, nil
}

func (r *MovieRepository) Find(ctx context.Context) ([]*model.Movie, error) {
	r.db = db.ConnectDB(r.dbCfg)
	m := query.Use(r.db).Movie

	movies, e := m.WithContext(ctx).Preload(m.Genres).Find()
	if e != nil {
		return nil, e
	}

	return movies, nil
}

func (r *MovieRepository) Save(ctx context.Context, target *NewMovie) (*model.Movie, error) {
	var err error
	defer func() {
		if err != nil {
			r.db.Rollback()
		} else {
			r.db.Commit()
		}
	}()

	r.db = db.ConnectDB(r.dbCfg)
	m := query.Use(r.db).Movie

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

func (r *MovieRepository) Delete(ctx context.Context, id *string) error {
	return nil
}
