package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	tests := []struct {
		id       string
		expected string
	}{
		{id: "1", expected: `{"title":"ターミネーター", "overview":"シュワちゃんカッコいい!!", "genres":["アクション", "ホラー", "SF"]}`},
		{id: "4", expected: `{"title":"ターミネーター4", "overview":"マーカス!!", "genres":["アクション", "ドラマ", "ホラー", "SF", "戦争"]}`},
	}

	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/movies/get/:id", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/movies/get/:id")
		c.SetParamNames("id")
		c.SetParamValues(tt.id)

		err := GetMovie(c)
		if err != nil {
			t.Fatal(err)
		}

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, tt.expected, rec.Body.String())
		}
	}
}
