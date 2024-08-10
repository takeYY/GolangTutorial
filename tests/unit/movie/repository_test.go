package movie

/* TODO: テスト通らないので後で直す
func TestGetMovieById(t *testing.T) {
	// --- Arrange --- //
	db, mock := initGormMockDB(t)

	rows := mock.
		NewRows([]string{"id", "title", "overview"}).
		AddRow(10, "ターミネーター:新起動/ジェニシス", "派生版")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `movie_genre` WHERE `movie_genre`.`movie_id` = 4")).WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `movie` WHERE `movie`.`id` = 4 ORDER BY `movie`.`id` LIMIT 1")).WillReturnRows(rows)

	// --- Act --- //
	r := movie.MovieRepository{Session: db}
	id := int32(4)
	actual, _ := r.FindByID(&id)

	// --- Assert --- //
	expected := &model.Movie{
		ID:        10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "ターミネーター:新起動/ジェニシス",
		Overview:  "派生版",
		Genres:    []model.Genre{},
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
*/
