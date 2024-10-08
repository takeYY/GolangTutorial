// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"golang_tutorial/model"
)

func newMovieGenre(db *gorm.DB, opts ...gen.DOOption) movieGenre {
	_movieGenre := movieGenre{}

	_movieGenre.movieGenreDo.UseDB(db, opts...)
	_movieGenre.movieGenreDo.UseModel(&model.MovieGenre{})

	tableName := _movieGenre.movieGenreDo.TableName()
	_movieGenre.ALL = field.NewAsterisk(tableName)
	_movieGenre.ID = field.NewInt32(tableName, "id")
	_movieGenre.CreatedAt = field.NewTime(tableName, "created_at")
	_movieGenre.UpdatedAt = field.NewTime(tableName, "updated_at")
	_movieGenre.MovieID = field.NewInt32(tableName, "movie_id")
	_movieGenre.GenreID = field.NewInt32(tableName, "genre_id")

	_movieGenre.fillFieldMap()

	return _movieGenre
}

type movieGenre struct {
	movieGenreDo

	ALL       field.Asterisk
	ID        field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	MovieID   field.Int32
	GenreID   field.Int32

	fieldMap map[string]field.Expr
}

func (m movieGenre) Table(newTableName string) *movieGenre {
	m.movieGenreDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m movieGenre) As(alias string) *movieGenre {
	m.movieGenreDo.DO = *(m.movieGenreDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *movieGenre) updateTableName(table string) *movieGenre {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.MovieID = field.NewInt32(table, "movie_id")
	m.GenreID = field.NewInt32(table, "genre_id")

	m.fillFieldMap()

	return m
}

func (m *movieGenre) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *movieGenre) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 5)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["movie_id"] = m.MovieID
	m.fieldMap["genre_id"] = m.GenreID
}

func (m movieGenre) clone(db *gorm.DB) movieGenre {
	m.movieGenreDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m movieGenre) replaceDB(db *gorm.DB) movieGenre {
	m.movieGenreDo.ReplaceDB(db)
	return m
}

type movieGenreDo struct{ gen.DO }

func (m movieGenreDo) Debug() *movieGenreDo {
	return m.withDO(m.DO.Debug())
}

func (m movieGenreDo) WithContext(ctx context.Context) *movieGenreDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m movieGenreDo) ReadDB() *movieGenreDo {
	return m.Clauses(dbresolver.Read)
}

func (m movieGenreDo) WriteDB() *movieGenreDo {
	return m.Clauses(dbresolver.Write)
}

func (m movieGenreDo) Session(config *gorm.Session) *movieGenreDo {
	return m.withDO(m.DO.Session(config))
}

func (m movieGenreDo) Clauses(conds ...clause.Expression) *movieGenreDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m movieGenreDo) Returning(value interface{}, columns ...string) *movieGenreDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m movieGenreDo) Not(conds ...gen.Condition) *movieGenreDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m movieGenreDo) Or(conds ...gen.Condition) *movieGenreDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m movieGenreDo) Select(conds ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m movieGenreDo) Where(conds ...gen.Condition) *movieGenreDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m movieGenreDo) Order(conds ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m movieGenreDo) Distinct(cols ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m movieGenreDo) Omit(cols ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m movieGenreDo) Join(table schema.Tabler, on ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m movieGenreDo) LeftJoin(table schema.Tabler, on ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m movieGenreDo) RightJoin(table schema.Tabler, on ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m movieGenreDo) Group(cols ...field.Expr) *movieGenreDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m movieGenreDo) Having(conds ...gen.Condition) *movieGenreDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m movieGenreDo) Limit(limit int) *movieGenreDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m movieGenreDo) Offset(offset int) *movieGenreDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m movieGenreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *movieGenreDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m movieGenreDo) Unscoped() *movieGenreDo {
	return m.withDO(m.DO.Unscoped())
}

func (m movieGenreDo) Create(values ...*model.MovieGenre) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m movieGenreDo) CreateInBatches(values []*model.MovieGenre, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m movieGenreDo) Save(values ...*model.MovieGenre) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m movieGenreDo) First() (*model.MovieGenre, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieGenre), nil
	}
}

func (m movieGenreDo) Take() (*model.MovieGenre, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieGenre), nil
	}
}

func (m movieGenreDo) Last() (*model.MovieGenre, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieGenre), nil
	}
}

func (m movieGenreDo) Find() ([]*model.MovieGenre, error) {
	result, err := m.DO.Find()
	return result.([]*model.MovieGenre), err
}

func (m movieGenreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MovieGenre, err error) {
	buf := make([]*model.MovieGenre, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m movieGenreDo) FindInBatches(result *[]*model.MovieGenre, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m movieGenreDo) Attrs(attrs ...field.AssignExpr) *movieGenreDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m movieGenreDo) Assign(attrs ...field.AssignExpr) *movieGenreDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m movieGenreDo) Joins(fields ...field.RelationField) *movieGenreDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m movieGenreDo) Preload(fields ...field.RelationField) *movieGenreDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m movieGenreDo) FirstOrInit() (*model.MovieGenre, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieGenre), nil
	}
}

func (m movieGenreDo) FirstOrCreate() (*model.MovieGenre, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieGenre), nil
	}
}

func (m movieGenreDo) FindByPage(offset int, limit int) (result []*model.MovieGenre, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m movieGenreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m movieGenreDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m movieGenreDo) Delete(models ...*model.MovieGenre) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *movieGenreDo) withDO(do gen.Dao) *movieGenreDo {
	m.DO = *do.(*gen.DO)
	return m
}
