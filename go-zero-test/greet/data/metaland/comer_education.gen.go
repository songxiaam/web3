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

func newComerEducation(db *gorm.DB, opts ...gen.DOOption) comerEducation {
	_comerEducation := comerEducation{}

	_comerEducation.comerEducationDo.UseDB(db, opts...)
	_comerEducation.comerEducationDo.UseModel(&model.ComerEducation{})

	tableName := _comerEducation.comerEducationDo.TableName()
	_comerEducation.ALL = field.NewAsterisk(tableName)
	_comerEducation.ID = field.NewInt64(tableName, "id")
	_comerEducation.ComerID = field.NewInt64(tableName, "comer_id")
	_comerEducation.School = field.NewString(tableName, "school")
	_comerEducation.Degree = field.NewString(tableName, "degree")
	_comerEducation.Major = field.NewString(tableName, "major")
	_comerEducation.StartDate = field.NewTime(tableName, "start_date")
	_comerEducation.EndDate = field.NewTime(tableName, "end_date")
	_comerEducation.Description = field.NewString(tableName, "description")
	_comerEducation.CreatedAt = field.NewTime(tableName, "created_at")
	_comerEducation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_comerEducation.IsDeleted = field.NewBool(tableName, "is_deleted")

	_comerEducation.fillFieldMap()

	return _comerEducation
}

type comerEducation struct {
	comerEducationDo

	ALL         field.Asterisk
	ID          field.Int64
	ComerID     field.Int64  // 用户ID
	School      field.String // 学校名称
	Degree      field.String // 学位
	Major       field.String // 专业
	StartDate   field.Time   // 开始日期
	EndDate     field.Time   // 结束日期
	Description field.String // 描述
	CreatedAt   field.Time
	UpdatedAt   field.Time
	IsDeleted   field.Bool // 是否删除

	fieldMap map[string]field.Expr
}

func (c comerEducation) Table(newTableName string) *comerEducation {
	c.comerEducationDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comerEducation) As(alias string) *comerEducation {
	c.comerEducationDo.DO = *(c.comerEducationDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comerEducation) updateTableName(table string) *comerEducation {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.ComerID = field.NewInt64(table, "comer_id")
	c.School = field.NewString(table, "school")
	c.Degree = field.NewString(table, "degree")
	c.Major = field.NewString(table, "major")
	c.StartDate = field.NewTime(table, "start_date")
	c.EndDate = field.NewTime(table, "end_date")
	c.Description = field.NewString(table, "description")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.IsDeleted = field.NewBool(table, "is_deleted")

	c.fillFieldMap()

	return c
}

func (c *comerEducation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comerEducation) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["id"] = c.ID
	c.fieldMap["comer_id"] = c.ComerID
	c.fieldMap["school"] = c.School
	c.fieldMap["degree"] = c.Degree
	c.fieldMap["major"] = c.Major
	c.fieldMap["start_date"] = c.StartDate
	c.fieldMap["end_date"] = c.EndDate
	c.fieldMap["description"] = c.Description
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["is_deleted"] = c.IsDeleted
}

func (c comerEducation) clone(db *gorm.DB) comerEducation {
	c.comerEducationDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comerEducation) replaceDB(db *gorm.DB) comerEducation {
	c.comerEducationDo.ReplaceDB(db)
	return c
}

type comerEducationDo struct{ gen.DO }

type IComerEducationDo interface {
	gen.SubQuery
	Debug() IComerEducationDo
	WithContext(ctx context.Context) IComerEducationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IComerEducationDo
	WriteDB() IComerEducationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IComerEducationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IComerEducationDo
	Not(conds ...gen.Condition) IComerEducationDo
	Or(conds ...gen.Condition) IComerEducationDo
	Select(conds ...field.Expr) IComerEducationDo
	Where(conds ...gen.Condition) IComerEducationDo
	Order(conds ...field.Expr) IComerEducationDo
	Distinct(cols ...field.Expr) IComerEducationDo
	Omit(cols ...field.Expr) IComerEducationDo
	Join(table schema.Tabler, on ...field.Expr) IComerEducationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IComerEducationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IComerEducationDo
	Group(cols ...field.Expr) IComerEducationDo
	Having(conds ...gen.Condition) IComerEducationDo
	Limit(limit int) IComerEducationDo
	Offset(offset int) IComerEducationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IComerEducationDo
	Unscoped() IComerEducationDo
	Create(values ...*model.ComerEducation) error
	CreateInBatches(values []*model.ComerEducation, batchSize int) error
	Save(values ...*model.ComerEducation) error
	First() (*model.ComerEducation, error)
	Take() (*model.ComerEducation, error)
	Last() (*model.ComerEducation, error)
	Find() ([]*model.ComerEducation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ComerEducation, err error)
	FindInBatches(result *[]*model.ComerEducation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ComerEducation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IComerEducationDo
	Assign(attrs ...field.AssignExpr) IComerEducationDo
	Joins(fields ...field.RelationField) IComerEducationDo
	Preload(fields ...field.RelationField) IComerEducationDo
	FirstOrInit() (*model.ComerEducation, error)
	FirstOrCreate() (*model.ComerEducation, error)
	FindByPage(offset int, limit int) (result []*model.ComerEducation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IComerEducationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c comerEducationDo) Debug() IComerEducationDo {
	return c.withDO(c.DO.Debug())
}

func (c comerEducationDo) WithContext(ctx context.Context) IComerEducationDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c comerEducationDo) ReadDB() IComerEducationDo {
	return c.Clauses(dbresolver.Read)
}

func (c comerEducationDo) WriteDB() IComerEducationDo {
	return c.Clauses(dbresolver.Write)
}

func (c comerEducationDo) Session(config *gorm.Session) IComerEducationDo {
	return c.withDO(c.DO.Session(config))
}

func (c comerEducationDo) Clauses(conds ...clause.Expression) IComerEducationDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c comerEducationDo) Returning(value interface{}, columns ...string) IComerEducationDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c comerEducationDo) Not(conds ...gen.Condition) IComerEducationDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c comerEducationDo) Or(conds ...gen.Condition) IComerEducationDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c comerEducationDo) Select(conds ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c comerEducationDo) Where(conds ...gen.Condition) IComerEducationDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c comerEducationDo) Order(conds ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c comerEducationDo) Distinct(cols ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c comerEducationDo) Omit(cols ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c comerEducationDo) Join(table schema.Tabler, on ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c comerEducationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c comerEducationDo) RightJoin(table schema.Tabler, on ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c comerEducationDo) Group(cols ...field.Expr) IComerEducationDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c comerEducationDo) Having(conds ...gen.Condition) IComerEducationDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c comerEducationDo) Limit(limit int) IComerEducationDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c comerEducationDo) Offset(offset int) IComerEducationDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c comerEducationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IComerEducationDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c comerEducationDo) Unscoped() IComerEducationDo {
	return c.withDO(c.DO.Unscoped())
}

func (c comerEducationDo) Create(values ...*model.ComerEducation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c comerEducationDo) CreateInBatches(values []*model.ComerEducation, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c comerEducationDo) Save(values ...*model.ComerEducation) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c comerEducationDo) First() (*model.ComerEducation, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerEducation), nil
	}
}

func (c comerEducationDo) Take() (*model.ComerEducation, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerEducation), nil
	}
}

func (c comerEducationDo) Last() (*model.ComerEducation, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerEducation), nil
	}
}

func (c comerEducationDo) Find() ([]*model.ComerEducation, error) {
	result, err := c.DO.Find()
	return result.([]*model.ComerEducation), err
}

func (c comerEducationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ComerEducation, err error) {
	buf := make([]*model.ComerEducation, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c comerEducationDo) FindInBatches(result *[]*model.ComerEducation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c comerEducationDo) Attrs(attrs ...field.AssignExpr) IComerEducationDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c comerEducationDo) Assign(attrs ...field.AssignExpr) IComerEducationDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c comerEducationDo) Joins(fields ...field.RelationField) IComerEducationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c comerEducationDo) Preload(fields ...field.RelationField) IComerEducationDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c comerEducationDo) FirstOrInit() (*model.ComerEducation, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerEducation), nil
	}
}

func (c comerEducationDo) FirstOrCreate() (*model.ComerEducation, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComerEducation), nil
	}
}

func (c comerEducationDo) FindByPage(offset int, limit int) (result []*model.ComerEducation, count int64, err error) {
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

func (c comerEducationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c comerEducationDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c comerEducationDo) Delete(models ...*model.ComerEducation) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *comerEducationDo) withDO(do gen.Dao) *comerEducationDo {
	c.DO = *do.(*gen.DO)
	return c
}
