package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// --- Arrange --- //
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// --- Act --- //
	c := e.NewContext(req, rec)

	// --- Assert --- //
	if err := Hello(c); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
