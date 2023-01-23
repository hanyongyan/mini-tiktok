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

	"awesomeProject4/dal/model"
)

func newTComment(db *gorm.DB, opts ...gen.DOOption) tComment {
	_tComment := tComment{}

	_tComment.tCommentDo.UseDB(db, opts...)
	_tComment.tCommentDo.UseModel(&model.TComment{})

	tableName := _tComment.tCommentDo.TableName()
	_tComment.ALL = field.NewAsterisk(tableName)
	_tComment.ID = field.NewInt64(tableName, "id")
	_tComment.UserID = field.NewInt64(tableName, "user_id")
	_tComment.Content = field.NewString(tableName, "content")
	_tComment.CreateDate = field.NewTime(tableName, "create_date")
	_tComment.VideoID = field.NewInt64(tableName, "video_id")

	_tComment.fillFieldMap()

	return _tComment
}

type tComment struct {
	tCommentDo tCommentDo

	ALL        field.Asterisk
	ID         field.Int64  // 评论id
	UserID     field.Int64  // 用户id
	Content    field.String // 评论内容
	CreateDate field.Time   // 评论发布日期，格式为mm-dd
	VideoID    field.Int64  // 视频id

	fieldMap map[string]field.Expr
}

func (t tComment) Table(newTableName string) *tComment {
	t.tCommentDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tComment) As(alias string) *tComment {
	t.tCommentDo.DO = *(t.tCommentDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tComment) updateTableName(table string) *tComment {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.UserID = field.NewInt64(table, "user_id")
	t.Content = field.NewString(table, "content")
	t.CreateDate = field.NewTime(table, "create_date")
	t.VideoID = field.NewInt64(table, "video_id")

	t.fillFieldMap()

	return t
}

func (t *tComment) WithContext(ctx context.Context) ITCommentDo { return t.tCommentDo.WithContext(ctx) }

func (t tComment) TableName() string { return t.tCommentDo.TableName() }

func (t tComment) Alias() string { return t.tCommentDo.Alias() }

func (t *tComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tComment) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["id"] = t.ID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["content"] = t.Content
	t.fieldMap["create_date"] = t.CreateDate
	t.fieldMap["video_id"] = t.VideoID
}

func (t tComment) clone(db *gorm.DB) tComment {
	t.tCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tComment) replaceDB(db *gorm.DB) tComment {
	t.tCommentDo.ReplaceDB(db)
	return t
}

type tCommentDo struct{ gen.DO }

type ITCommentDo interface {
	gen.SubQuery
	Debug() ITCommentDo
	WithContext(ctx context.Context) ITCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITCommentDo
	WriteDB() ITCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITCommentDo
	Not(conds ...gen.Condition) ITCommentDo
	Or(conds ...gen.Condition) ITCommentDo
	Select(conds ...field.Expr) ITCommentDo
	Where(conds ...gen.Condition) ITCommentDo
	Order(conds ...field.Expr) ITCommentDo
	Distinct(cols ...field.Expr) ITCommentDo
	Omit(cols ...field.Expr) ITCommentDo
	Join(table schema.Tabler, on ...field.Expr) ITCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITCommentDo
	Group(cols ...field.Expr) ITCommentDo
	Having(conds ...gen.Condition) ITCommentDo
	Limit(limit int) ITCommentDo
	Offset(offset int) ITCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITCommentDo
	Unscoped() ITCommentDo
	Create(values ...*model.TComment) error
	CreateInBatches(values []*model.TComment, batchSize int) error
	Save(values ...*model.TComment) error
	First() (*model.TComment, error)
	Take() (*model.TComment, error)
	Last() (*model.TComment, error)
	Find() ([]*model.TComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TComment, err error)
	FindInBatches(result *[]*model.TComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITCommentDo
	Assign(attrs ...field.AssignExpr) ITCommentDo
	Joins(fields ...field.RelationField) ITCommentDo
	Preload(fields ...field.RelationField) ITCommentDo
	FirstOrInit() (*model.TComment, error)
	FirstOrCreate() (*model.TComment, error)
	FindByPage(offset int, limit int) (result []*model.TComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tCommentDo) Debug() ITCommentDo {
	return t.withDO(t.DO.Debug())
}

func (t tCommentDo) WithContext(ctx context.Context) ITCommentDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tCommentDo) ReadDB() ITCommentDo {
	return t.Clauses(dbresolver.Read)
}

func (t tCommentDo) WriteDB() ITCommentDo {
	return t.Clauses(dbresolver.Write)
}

func (t tCommentDo) Session(config *gorm.Session) ITCommentDo {
	return t.withDO(t.DO.Session(config))
}

func (t tCommentDo) Clauses(conds ...clause.Expression) ITCommentDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tCommentDo) Returning(value interface{}, columns ...string) ITCommentDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tCommentDo) Not(conds ...gen.Condition) ITCommentDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tCommentDo) Or(conds ...gen.Condition) ITCommentDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tCommentDo) Select(conds ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tCommentDo) Where(conds ...gen.Condition) ITCommentDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITCommentDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tCommentDo) Order(conds ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tCommentDo) Distinct(cols ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tCommentDo) Omit(cols ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tCommentDo) Join(table schema.Tabler, on ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tCommentDo) Group(cols ...field.Expr) ITCommentDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tCommentDo) Having(conds ...gen.Condition) ITCommentDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tCommentDo) Limit(limit int) ITCommentDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tCommentDo) Offset(offset int) ITCommentDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITCommentDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tCommentDo) Unscoped() ITCommentDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tCommentDo) Create(values ...*model.TComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tCommentDo) CreateInBatches(values []*model.TComment, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tCommentDo) Save(values ...*model.TComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tCommentDo) First() (*model.TComment, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TComment), nil
	}
}

func (t tCommentDo) Take() (*model.TComment, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TComment), nil
	}
}

func (t tCommentDo) Last() (*model.TComment, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TComment), nil
	}
}

func (t tCommentDo) Find() ([]*model.TComment, error) {
	result, err := t.DO.Find()
	return result.([]*model.TComment), err
}

func (t tCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TComment, err error) {
	buf := make([]*model.TComment, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tCommentDo) FindInBatches(result *[]*model.TComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tCommentDo) Attrs(attrs ...field.AssignExpr) ITCommentDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tCommentDo) Assign(attrs ...field.AssignExpr) ITCommentDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tCommentDo) Joins(fields ...field.RelationField) ITCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tCommentDo) Preload(fields ...field.RelationField) ITCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tCommentDo) FirstOrInit() (*model.TComment, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TComment), nil
	}
}

func (t tCommentDo) FirstOrCreate() (*model.TComment, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TComment), nil
	}
}

func (t tCommentDo) FindByPage(offset int, limit int) (result []*model.TComment, count int64, err error) {
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

func (t tCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tCommentDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tCommentDo) Delete(models ...*model.TComment) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tCommentDo) withDO(do gen.Dao) *tCommentDo {
	t.DO = *do.(*gen.DO)
	return t
}