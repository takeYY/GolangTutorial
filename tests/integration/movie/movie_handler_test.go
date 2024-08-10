package movie

import (
	"encoding/json"
	"golang_tutorial/config"
	"golang_tutorial/internal/movie"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	tests := []struct {
		id       string
		expected movie.MovieResponse
	}{
		{
			id: "1",
			expected: movie.MovieResponse{
				ID:       1,
				Title:    "ターミネーター",
				Overview: "シュワちゃんカッコいい!!",
				Genres: []movie.Genre{
					{
						Code: "action",
						Name: "アクション",
					},
					{
						Code: "horror",
						Name: "ホラー",
					},
					{
						Code: "sf",
						Name: "SF",
					},
				},
			},
		},
		{
			id: "4",
			expected: movie.MovieResponse{
				ID:       4,
				Title:    "ターミネーター4",
				Overview: "マーカス!!",
				Genres: []movie.Genre{
					{
						Code: "action",
						Name: "アクション",
					},
					{
						Code: "drama",
						Name: "ドラマ",
					},
					{
						Code: "horror",
						Name: "ホラー",
					},
					{
						Code: "sf",
						Name: "SF",
					},
					{
						Code: "war",
						Name: "戦争",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/movies/:id", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/movies/:id")
		c.SetParamNames("id")
		c.SetParamValues(tt.id)

		// 設定ファイルの読み込み
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		target := movie.NewHandler(&cfg.ReaderDatabase, &cfg.WriterDatabase)
		er := target.ReaderHandler.GetMovie(c)
		if er != nil {
			t.Fatal(er)
		}

		if assert.NoError(t, er) {
			assert.Equal(t, http.StatusOK, rec.Code)

			jsonBytes, _ := json.Marshal(tt.expected)
			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}
	}
}
