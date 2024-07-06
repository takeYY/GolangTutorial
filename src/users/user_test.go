package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)



func TestGetUserName(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	c := e.NewContext(req, rec)

	// NOTE: 関数のテストではなく、サーバを立ててテストした方が良いかも
	err := c.JSON(http.StatusOK, User{Name: "John Doe", Description: "説明のテストですよ"})
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"name":"John Doe", "description":"説明のテストですよ"}`, rec.Body.String())
	}
}
