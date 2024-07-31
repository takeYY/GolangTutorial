// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Genre:      newGenre(db, opts...),
		GenreNeo:   newGenreNeo(db, opts...),
		Movie:      newMovie(db, opts...),
		MovieGenre: newMovieGenre(db, opts...),
		MovieNeo:   newMovieNeo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Genre      genre
	GenreNeo   genreNeo
	Movie      movie
	MovieGenre movieGenre
	MovieNeo   movieNeo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Genre:      q.Genre.clone(db),
		GenreNeo:   q.GenreNeo.clone(db),
		Movie:      q.Movie.clone(db),
		MovieGenre: q.MovieGenre.clone(db),
		MovieNeo:   q.MovieNeo.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Genre:      q.Genre.replaceDB(db),
		GenreNeo:   q.GenreNeo.replaceDB(db),
		Movie:      q.Movie.replaceDB(db),
		MovieGenre: q.MovieGenre.replaceDB(db),
		MovieNeo:   q.MovieNeo.replaceDB(db),
	}
}

type queryCtx struct {
	Genre      *genreDo
	GenreNeo   *genreNeoDo
	Movie      *movieDo
	MovieGenre *movieGenreDo
	MovieNeo   *movieNeoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Genre:      q.Genre.WithContext(ctx),
		GenreNeo:   q.GenreNeo.WithContext(ctx),
		Movie:      q.Movie.WithContext(ctx),
		MovieGenre: q.MovieGenre.WithContext(ctx),
		MovieNeo:   q.MovieNeo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
