package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUserName(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{name: "John Doe", expected: `{"name":"John Doe", "description":"テストの説明_1"}`},
		{name: "Jane Doe", expected: `{"name":"Jane Doe", "description":"説明_2"}`},
	}

	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/users/:name")
		c.SetParamNames("name")
		c.SetParamValues(tt.name)

		err := GetUserName(c)
		if err != nil {
			t.Fatal(err)
		}

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, tt.expected, rec.Body.String())
		}
	}
}

func TestGetUserNameNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:name")
	c.SetParamNames("name")
	c.SetParamValues("invalid user")

	err := GetUserName(c)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.JSONEq(t, `{"message":"invalid user is not found."}`, rec.Body.String())
}
