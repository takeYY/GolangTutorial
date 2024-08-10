package movie

import (
	"context"
	"golang_tutorial/config"
)

type CommandService struct {
	MovieRepo ICommandRepository
}

func NewCommandService(cfg *config.Config) *CommandService {
	var movieRepo ICommandRepository = &MovieRepository{
		dbCfg: &cfg.WriterDatabase,
	}

	return &CommandService{
		MovieRepo: movieRepo,
	}
}

func (cs *CommandService) CreateMovie(ctx context.Context, movie *NewMovie) (*MovieResponse, error) {
	result, err := cs.MovieRepo.Save(ctx, movie)
	if err != nil {
		return nil, err
	}

	movieResponse := &MovieResponse{
		ID:       result.ID,
		Title:    result.Title,
		Overview: result.Overview,
		Genres:   []Genre{},
	}

	return movieResponse, nil
}
