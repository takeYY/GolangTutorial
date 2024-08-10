package movie

import (
	"context"
	"golang_tutorial/config"
)

type QueryService struct {
	MovieRepo IQueryRepository
}

// 新しいQueryServiceを作成
func NewQueryService(cfg *config.Config) *QueryService {
	var movieRepo IQueryRepository = &MovieRepository{
		dbCfg: &cfg.ReaderDatabase,
	}

	return &QueryService{
		MovieRepo: movieRepo,
	}
}

// 指定された映画IDの詳細情報を取得
func (qs *QueryService) GetMovieDetails(ctx context.Context, movieID *int32) (*MovieResponse, error) {
	movie, err := qs.MovieRepo.FindByID(ctx, movieID)
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

func (qs *QueryService) GetMovies(ctx context.Context) ([]MovieResponse, error) {
	result, err := qs.MovieRepo.Find(ctx)
	if err != nil {
		return nil, err
	}

	var movies []MovieResponse
	for _, res := range result {
		var genres []Genre
		for _, genre := range res.Genres {
			g := Genre{
				Code: genre.Code,
				Name: genre.Name,
			}

			genres = append(genres, g)
		}

		m := MovieResponse{
			ID:       res.ID,
			Title:    res.Title,
			Overview: res.Overview,
			Genres:   genres,
		}
		movies = append(movies, m)
	}

	return movies, nil
}
