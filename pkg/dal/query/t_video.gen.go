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

	"mini_tiktok/pkg/dal/model"
)

func newTVideo(db *gorm.DB, opts ...gen.DOOption) tVideo {
	_tVideo := tVideo{}

	_tVideo.tVideoDo.UseDB(db, opts...)
	_tVideo.tVideoDo.UseModel(&model.TVideo{})

	tableName := _tVideo.tVideoDo.TableName()
	_tVideo.ALL = field.NewAsterisk(tableName)
	_tVideo.ID = field.NewInt64(tableName, "id")
	_tVideo.AuthorID = field.NewInt64(tableName, "author_id")
	_tVideo.PlayURL = field.NewString(tableName, "play_url")
	_tVideo.CoverURL = field.NewString(tableName, "cover_url")
	_tVideo.FavoriteCount = field.NewInt64(tableName, "favorite_count")
	_tVideo.CommentCount = field.NewInt64(tableName, "comment_count")
	_tVideo.IsFavorite = field.NewBool(tableName, "is_favorite")
	_tVideo.Title = field.NewString(tableName, "title")

	_tVideo.fillFieldMap()

	return _tVideo
}

type tVideo struct {
	tVideoDo tVideoDo

	ALL           field.Asterisk
	ID            field.Int64  // 视频id
	AuthorID      field.Int64  // 作者id
	PlayURL       field.String // 视频链接
	CoverURL      field.String // 视频封面链接
	FavoriteCount field.Int64  // 点赞数
	CommentCount  field.Int64  // 评论数
	IsFavorite    field.Bool   // 是否已点赞(0为未点赞, 1为已点赞)
	Title         field.String // 视频标题

	fieldMap map[string]field.Expr
}

func (t tVideo) Table(newTableName string) *tVideo {
	t.tVideoDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tVideo) As(alias string) *tVideo {
	t.tVideoDo.DO = *(t.tVideoDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tVideo) updateTableName(table string) *tVideo {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.AuthorID = field.NewInt64(table, "author_id")
	t.PlayURL = field.NewString(table, "play_url")
	t.CoverURL = field.NewString(table, "cover_url")
	t.FavoriteCount = field.NewInt64(table, "favorite_count")
	t.CommentCount = field.NewInt64(table, "comment_count")
	t.IsFavorite = field.NewBool(table, "is_favorite")
	t.Title = field.NewString(table, "title")

	t.fillFieldMap()

	return t
}

func (t *tVideo) WithContext(ctx context.Context) ITVideoDo { return t.tVideoDo.WithContext(ctx) }

func (t tVideo) TableName() string { return t.tVideoDo.TableName() }

func (t tVideo) Alias() string { return t.tVideoDo.Alias() }

func (t *tVideo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tVideo) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.ID
	t.fieldMap["author_id"] = t.AuthorID
	t.fieldMap["play_url"] = t.PlayURL
	t.fieldMap["cover_url"] = t.CoverURL
	t.fieldMap["favorite_count"] = t.FavoriteCount
	t.fieldMap["comment_count"] = t.CommentCount
	t.fieldMap["is_favorite"] = t.IsFavorite
	t.fieldMap["title"] = t.Title
}

func (t tVideo) clone(db *gorm.DB) tVideo {
	t.tVideoDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tVideo) replaceDB(db *gorm.DB) tVideo {
	t.tVideoDo.ReplaceDB(db)
	return t
}

type tVideoDo struct{ gen.DO }

type ITVideoDo interface {
	gen.SubQuery
	Debug() ITVideoDo
	WithContext(ctx context.Context) ITVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITVideoDo
	WriteDB() ITVideoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITVideoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITVideoDo
	Not(conds ...gen.Condition) ITVideoDo
	Or(conds ...gen.Condition) ITVideoDo
	Select(conds ...field.Expr) ITVideoDo
	Where(conds ...gen.Condition) ITVideoDo
	Order(conds ...field.Expr) ITVideoDo
	Distinct(cols ...field.Expr) ITVideoDo
	Omit(cols ...field.Expr) ITVideoDo
	Join(table schema.Tabler, on ...field.Expr) ITVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITVideoDo
	Group(cols ...field.Expr) ITVideoDo
	Having(conds ...gen.Condition) ITVideoDo
	Limit(limit int) ITVideoDo
	Offset(offset int) ITVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITVideoDo
	Unscoped() ITVideoDo
	Create(values ...*model.TVideo) error
	CreateInBatches(values []*model.TVideo, batchSize int) error
	Save(values ...*model.TVideo) error
	First() (*model.TVideo, error)
	Take() (*model.TVideo, error)
	Last() (*model.TVideo, error)
	Find() ([]*model.TVideo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TVideo, err error)
	FindInBatches(result *[]*model.TVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TVideo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITVideoDo
	Assign(attrs ...field.AssignExpr) ITVideoDo
	Joins(fields ...field.RelationField) ITVideoDo
	Preload(fields ...field.RelationField) ITVideoDo
	FirstOrInit() (*model.TVideo, error)
	FirstOrCreate() (*model.TVideo, error)
	FindByPage(offset int, limit int) (result []*model.TVideo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tVideoDo) Debug() ITVideoDo {
	return t.withDO(t.DO.Debug())
}

func (t tVideoDo) WithContext(ctx context.Context) ITVideoDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tVideoDo) ReadDB() ITVideoDo {
	return t.Clauses(dbresolver.Read)
}

func (t tVideoDo) WriteDB() ITVideoDo {
	return t.Clauses(dbresolver.Write)
}

func (t tVideoDo) Session(config *gorm.Session) ITVideoDo {
	return t.withDO(t.DO.Session(config))
}

func (t tVideoDo) Clauses(conds ...clause.Expression) ITVideoDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tVideoDo) Returning(value interface{}, columns ...string) ITVideoDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tVideoDo) Not(conds ...gen.Condition) ITVideoDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tVideoDo) Or(conds ...gen.Condition) ITVideoDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tVideoDo) Select(conds ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tVideoDo) Where(conds ...gen.Condition) ITVideoDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tVideoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITVideoDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tVideoDo) Order(conds ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tVideoDo) Distinct(cols ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tVideoDo) Omit(cols ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tVideoDo) Join(table schema.Tabler, on ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tVideoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tVideoDo) RightJoin(table schema.Tabler, on ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tVideoDo) Group(cols ...field.Expr) ITVideoDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tVideoDo) Having(conds ...gen.Condition) ITVideoDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tVideoDo) Limit(limit int) ITVideoDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tVideoDo) Offset(offset int) ITVideoDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tVideoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITVideoDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tVideoDo) Unscoped() ITVideoDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tVideoDo) Create(values ...*model.TVideo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tVideoDo) CreateInBatches(values []*model.TVideo, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tVideoDo) Save(values ...*model.TVideo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tVideoDo) First() (*model.TVideo, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVideo), nil
	}
}

func (t tVideoDo) Take() (*model.TVideo, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVideo), nil
	}
}

func (t tVideoDo) Last() (*model.TVideo, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVideo), nil
	}
}

func (t tVideoDo) Find() ([]*model.TVideo, error) {
	result, err := t.DO.Find()
	return result.([]*model.TVideo), err
}

func (t tVideoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TVideo, err error) {
	buf := make([]*model.TVideo, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tVideoDo) FindInBatches(result *[]*model.TVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tVideoDo) Attrs(attrs ...field.AssignExpr) ITVideoDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tVideoDo) Assign(attrs ...field.AssignExpr) ITVideoDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tVideoDo) Joins(fields ...field.RelationField) ITVideoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tVideoDo) Preload(fields ...field.RelationField) ITVideoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tVideoDo) FirstOrInit() (*model.TVideo, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVideo), nil
	}
}

func (t tVideoDo) FirstOrCreate() (*model.TVideo, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TVideo), nil
	}
}

func (t tVideoDo) FindByPage(offset int, limit int) (result []*model.TVideo, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tVideoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tVideoDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tVideoDo) Delete(models ...*model.TVideo) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tVideoDo) withDO(do gen.Dao) *tVideoDo {
	t.DO = *do.(*gen.DO)
	return t
}
