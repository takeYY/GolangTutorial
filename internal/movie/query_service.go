package movie

import (
	"golang_tutorial/config"
	"golang_tutorial/db"
)

type QueryService struct {
	MovieRepo IRepository
}

// 新しいQueryServiceを作成
func NewQueryService(dbCfg *config.Database) *QueryService {
	readerDBSession := db.ConnectDB(dbCfg)
	readerTx := readerDBSession.Begin()

	var movieRepo IRepository = &MovieRepository{
		Session: readerTx,
	}

	return &QueryService{
		MovieRepo: movieRepo,
	}
}

// 指定された映画IDの詳細情報を取得
func (qs *QueryService) GetMovieDetails(movieID *int32) (*MovieResponse, error) {
	movie, err := qs.MovieRepo.FindByID(movieID)
	if err != nil {
		return nil, err
	}

	var genres []Genre
	for _, genre := range movie.Genres {
		var g Genre
		g.Code = genre.Code
		g.Name = genre.Name

		genres = append(genres, g)
	}

	movieResponse := &MovieResponse{
		ID:       movie.ID,
		Title:    movie.Title,
		Overview: movie.Overview,
		Genres:   genres,
	}

	return movieResponse, nil
}
