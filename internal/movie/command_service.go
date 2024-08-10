package movie

import (
	"golang_tutorial/config"
	"golang_tutorial/db"
)

type CommandService struct {
	MovieRepo ICommandRepository
}

func NewCommandService(dbCfg *config.Database) *CommandService {
	writerDBSession := db.ConnectDB(dbCfg)
	writerTx := writerDBSession.Begin()

	var movieRepo ICommandRepository = &MovieRepository{
		Session: writerTx,
	}

	return &CommandService{
		MovieRepo: movieRepo,
	}
}

func (cs *CommandService) CreateMovie(movie *NewMovie) (*MovieResponse, error) {
	result, err := cs.MovieRepo.Save(movie)
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
