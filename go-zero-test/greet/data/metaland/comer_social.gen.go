// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package metaland

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"greet/data/model"
)

func newComerSocial(db *gorm.DB, opts ...gen.DOOption) comerSocial {
	_comerSocial := comerSocial{}

	_comerSocial.comerSocialDo.UseDB(db, opts...)
	_comerSocial.comerSocialDo.UseModel(&model.ComerSocial{})

	tableName := _comerSocial.comerSocialDo.TableName()
	_comerSocial.ALL = field.NewAsterisk(tableName)
	_comerSocial.ID = field.NewInt64(tableName, "id")
	_comerSocial.ComerID = field.NewInt64(tableName, "comer_id")
	_comerSocial.Platform = field.NewString(tableName, "platform")
	_comerSocial.Username = field.NewString(tableName, "username")
	_comerSocial.URL = field.NewString(tableName, "url")
	_comerSocial.IsVerified = field.NewBool(tableName, "is_verified")
	_comerSocial.CreatedAt = field.NewTime(tableName, "created_at")
	_comerSocial.UpdatedAt = field.NewTime(tableName, "updated_at")
	_comerSocial.IsDeleted = field.NewBool(tableName, "is_deleted")

	_comerSocial.fillFieldMap()

	return _comerSocial
}

type comerSocial struct {
	comerSocialDo

	ALL        field.Asterisk
	ID         field.Int64
	ComerID    field.Int64  // 用户ID
	Platform   field.String // 平台(twitter/discord/telegram等)
	Username   field.String // 用户名
	URL        field.String // 链接
	IsVerified field.Bool   // 是否认证
	CreatedAt  field.Time
	UpdatedAt  field.Time
	IsDeleted  field.Bool // 是否删除

	fieldMap map[string]field.Expr
}

func (c comerSocial) Table(newTableName string) *comerSocial {
	c.comerSocialDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comerSocial) As(alias string) *comerSocial {
	c.comerSocialDo.DO = *(c.comerSocialDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comerSocial) updateTableName(table string) *comerSocial {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.ComerID = field.NewInt64(table, "comer_id")
	c.Platform = field.NewString(table, "platform")
	c.Username = field.NewString(table, "username")
	c.URL = field.NewString(table, "url")
	c.IsVerified = field.NewBool(table, "is_verified")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.IsDeleted = field.NewBool(table, "is_deleted")

	c.fillFieldMap()

	return c
}

func (c *comerSocial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comerSocial) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["comer_id"] = c.ComerID
	c.fieldMap["platform"] = c.Platform
	c.fieldMap["username"] = c.Username
	c.fieldMap["url"] = c.URL
	c.fieldMap["is_verified"] = c.IsVerified
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["is_deleted"] = c.IsDeleted
}

func (c comerSocial) clone(db *gorm.DB) comerSocial {
	c.comerSocialDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comerSocial) replaceDB(db *gorm.DB) comerSocial {
	c.comerSocialDo.ReplaceDB(db)
	return c
}

type comerSocialDo struct{ gen.DO }

type IComerSocialDo interface {
	gen.SubQuery
	Debug() IComerSocialDo
	WithContext(ctx context.Context) IComerSocialDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IComerSocialDo
	WriteDB() IComerSocialDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IComerSocialDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IComerSocialDo
	Not(conds ...gen.Condition) IComerSocialDo
	Or(conds ...gen.Condition) IComerSocialDo
	Select(conds ...field.Expr) IComerSocialDo
	Where(conds ...gen.Condition) IComerSocialDo
	Order(conds ...field.Expr) IComerSocialDo
	Distinct(cols ...field.Expr) IComerSocialDo
	Omit(cols ...field.Expr) IComerSocialDo
	Join(table schema.Tabler, on ...field.Expr) IComerSocialDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IComerSocialDo
	RightJoin(table schema.Tabler, on ...field.Expr) IComerSocialDo
	Group(cols ...field.Expr) IComerSocialDo
	Having(conds ...gen.Condition) IComerSocialDo
	Limit(limit int) IComerSocialDo
	Offset(offset int) IComerSocialDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IComerSocialDo
	Unscoped() IComerSocialDo
	Create(values ...*model.ComerSocial) error
	CreateInBatches(values []*model.ComerSocial, batchSize int) error
	Save(values ...*model.ComerSocial) error
	First() (*model.ComerSocial, error)
	Take() (*model.ComerSocial, error)
	Last() (*model.ComerSocial, error)
	Find() ([]*model.ComerSocial, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ComerSocial, err error)
	FindInBatches(result *[]*model.ComerSocial, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ComerSocial) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IComerSocialDo
	Assign(attrs ...field.AssignExpr) IComerSocialDo
	Joins(fields ...field.RelationField) IComerSocialDo
	Preload(fields ...field.RelationField) IComerSocialDo
	FirstOrInit() (*model.ComerSocial, error)
	FirstOrCreate() (*model.ComerSocial, error)
	FindByPage(offset int, limit int) (result []*model.ComerSocial, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IComerSocialDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c comerSocialDo) Debug() IComerSocialDo {
	return c.withDO(c.DO.Debug())
}

func (c comerSocialDo) WithContext(ctx context.Context) IComerSocialDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c comerSocialDo) ReadDB() IComerSocialDo {
	return c.Clauses(dbresolver.Read)
}

func (c comerSocialDo) WriteDB() IComerSocialDo {
	return c.Clauses(dbresolver.Write)
}

func (c comerSocialDo) Session(config *gorm.Session) IComerSocialDo {
	return c.withDO(c.DO.Session(config))
}

func (c comerSocialDo) Clauses(conds ...clause.Expression) IComerSocialDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c comerSocialDo) Returning(value interface{}, columns ...string) IComerSocialDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c comerSocialDo) Not(conds ...gen.Condition) IComerSocialDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c comerSocialDo) Or(conds ...gen.Condition) IComerSocialDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c comerSocialDo) Select(conds ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c comerSocialDo) Where(conds ...gen.Condition) IComerSocialDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c comerSocialDo) Order(conds ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c comerSocialDo) Distinct(cols ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c comerSocialDo) Omit(cols ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c comerSocialDo) Join(table schema.Tabler, on ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c comerSocialDo) LeftJoin(table schema.Tabler, on ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c comerSocialDo) RightJoin(table schema.Tabler, on ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c comerSocialDo) Group(cols ...field.Expr) IComerSocialDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c comerSocialDo) Having(conds ...gen.Condition) IComerSocialDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c comerSocialDo) Limit(limit int) IComerSocialDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c comerSocialDo) Offset(offset int) IComerSocialDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c comerSocialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IComerSocialDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c comerSocialDo) Unscoped() IComerSocialDo {
	return c.withDO(c.DO.Unscoped())
}

func (c comerSocialDo) Create(values ...*model.ComerSocial) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c comerSocialDo) CreateInBatches(values []*model.ComerSocial, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c comerSocialDo) Save(values ...*model.ComerSocial) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c comerSocialDo) First() (*model.ComerSocial, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerSocial), nil
	}
}

func (c comerSocialDo) Take() (*model.ComerSocial, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerSocial), nil
	}
}

func (c comerSocialDo) Last() (*model.ComerSocial, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerSocial), nil
	}
}

func (c comerSocialDo) Find() ([]*model.ComerSocial, error) {
	result, err := c.DO.Find()
	return result.([]*model.ComerSocial), err
}

func (c comerSocialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ComerSocial, err error) {
	buf := make([]*model.ComerSocial, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c comerSocialDo) FindInBatches(result *[]*model.ComerSocial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c comerSocialDo) Attrs(attrs ...field.AssignExpr) IComerSocialDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c comerSocialDo) Assign(attrs ...field.AssignExpr) IComerSocialDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c comerSocialDo) Joins(fields ...field.RelationField) IComerSocialDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c comerSocialDo) Preload(fields ...field.RelationField) IComerSocialDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c comerSocialDo) FirstOrInit() (*model.ComerSocial, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerSocial), nil
	}
}

func (c comerSocialDo) FirstOrCreate() (*model.ComerSocial, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerSocial), nil
	}
}

func (c comerSocialDo) FindByPage(offset int, limit int) (result []*model.ComerSocial, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c comerSocialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c comerSocialDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c comerSocialDo) Delete(models ...*model.ComerSocial) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *comerSocialDo) withDO(do gen.Dao) *comerSocialDo {
	c.DO = *do.(*gen.DO)
	return c
}
