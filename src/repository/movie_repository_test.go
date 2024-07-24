package repository

import (
	"golang_tutorial/src/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestGetMovies(t *testing.T) {
	// --- Arrange --- //
	db, mock := initGormMockDB(t)

	rows := mock.
		NewRows([]string{"id", "title", "overview"}).
		AddRow(10, "ターミネーター:新起動/ジェニシス", "派生版").
		AddRow(11, "ターミネーター:ニュー・フェイト", "What r u doing?")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `movie`")).WillReturnRows(rows)

	// --- Act --- //
	r := MovieRepository{Session: db}
	actual, _ := r.GetMovies()

	// --- Assert --- //
	expected := []*model.Movie{
		{
			ID:       10,
			Title:    "ターミネーター:新起動/ジェニシス",
			Overview: "派生版",
		},
		{
			ID:       11,
			Title:    "ターミネーター:ニュー・フェイト",
			Overview: "What r u doing?",
		},
	}
	assert.Equal(t, expected, actual)
}

func initGormMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		t.Fatalf("failed to use mocked gorm '%s'", err)
	}

	return mockDB, mock
}
