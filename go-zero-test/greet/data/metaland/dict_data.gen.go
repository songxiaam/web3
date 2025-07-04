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

func newDictDatum(db *gorm.DB, opts ...gen.DOOption) dictDatum {
	_dictDatum := dictDatum{}

	_dictDatum.dictDatumDo.UseDB(db, opts...)
	_dictDatum.dictDatumDo.UseModel(&model.DictDatum{})

	tableName := _dictDatum.dictDatumDo.TableName()
	_dictDatum.ALL = field.NewAsterisk(tableName)
	_dictDatum.ID = field.NewInt32(tableName, "id")
	_dictDatum.StartupID = field.NewInt64(tableName, "startup_id")
	_dictDatum.DictType = field.NewString(tableName, "dict_type")
	_dictDatum.DictLabel = field.NewString(tableName, "dict_label")
	_dictDatum.DictValue = field.NewString(tableName, "dict_value")
	_dictDatum.SeqNum = field.NewInt32(tableName, "seq_num")
	_dictDatum.Status = field.NewBool(tableName, "status")
	_dictDatum.Remark = field.NewString(tableName, "remark")
	_dictDatum.CreatedAt = field.NewTime(tableName, "created_at")
	_dictDatum.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dictDatum.IsDeleted = field.NewBool(tableName, "is_deleted")

	_dictDatum.fillFieldMap()

	return _dictDatum
}

type dictDatum struct {
	dictDatumDo

	ALL       field.Asterisk
	ID        field.Int32
	StartupID field.Int64
	DictType  field.String
	DictLabel field.String
	DictValue field.String
	SeqNum    field.Int32
	Status    field.Bool // 1:enabled 2:disabled
	Remark    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	IsDeleted field.Bool

	fieldMap map[string]field.Expr
}

func (d dictDatum) Table(newTableName string) *dictDatum {
	d.dictDatumDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dictDatum) As(alias string) *dictDatum {
	d.dictDatumDo.DO = *(d.dictDatumDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dictDatum) updateTableName(table string) *dictDatum {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.StartupID = field.NewInt64(table, "startup_id")
	d.DictType = field.NewString(table, "dict_type")
	d.DictLabel = field.NewString(table, "dict_label")
	d.DictValue = field.NewString(table, "dict_value")
	d.SeqNum = field.NewInt32(table, "seq_num")
	d.Status = field.NewBool(table, "status")
	d.Remark = field.NewString(table, "remark")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.IsDeleted = field.NewBool(table, "is_deleted")

	d.fillFieldMap()

	return d
}

func (d *dictDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dictDatum) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["id"] = d.ID
	d.fieldMap["startup_id"] = d.StartupID
	d.fieldMap["dict_type"] = d.DictType
	d.fieldMap["dict_label"] = d.DictLabel
	d.fieldMap["dict_value"] = d.DictValue
	d.fieldMap["seq_num"] = d.SeqNum
	d.fieldMap["status"] = d.Status
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["is_deleted"] = d.IsDeleted
}

func (d dictDatum) clone(db *gorm.DB) dictDatum {
	d.dictDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dictDatum) replaceDB(db *gorm.DB) dictDatum {
	d.dictDatumDo.ReplaceDB(db)
	return d
}

type dictDatumDo struct{ gen.DO }

type IDictDatumDo interface {
	gen.SubQuery
	Debug() IDictDatumDo
	WithContext(ctx context.Context) IDictDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDictDatumDo
	WriteDB() IDictDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDictDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDictDatumDo
	Not(conds ...gen.Condition) IDictDatumDo
	Or(conds ...gen.Condition) IDictDatumDo
	Select(conds ...field.Expr) IDictDatumDo
	Where(conds ...gen.Condition) IDictDatumDo
	Order(conds ...field.Expr) IDictDatumDo
	Distinct(cols ...field.Expr) IDictDatumDo
	Omit(cols ...field.Expr) IDictDatumDo
	Join(table schema.Tabler, on ...field.Expr) IDictDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDictDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDictDatumDo
	Group(cols ...field.Expr) IDictDatumDo
	Having(conds ...gen.Condition) IDictDatumDo
	Limit(limit int) IDictDatumDo
	Offset(offset int) IDictDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDictDatumDo
	Unscoped() IDictDatumDo
	Create(values ...*model.DictDatum) error
	CreateInBatches(values []*model.DictDatum, batchSize int) error
	Save(values ...*model.DictDatum) error
	First() (*model.DictDatum, error)
	Take() (*model.DictDatum, error)
	Last() (*model.DictDatum, error)
	Find() ([]*model.DictDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictDatum, err error)
	FindInBatches(result *[]*model.DictDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DictDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDictDatumDo
	Assign(attrs ...field.AssignExpr) IDictDatumDo
	Joins(fields ...field.RelationField) IDictDatumDo
	Preload(fields ...field.RelationField) IDictDatumDo
	FirstOrInit() (*model.DictDatum, error)
	FirstOrCreate() (*model.DictDatum, error)
	FindByPage(offset int, limit int) (result []*model.DictDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDictDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dictDatumDo) Debug() IDictDatumDo {
	return d.withDO(d.DO.Debug())
}

func (d dictDatumDo) WithContext(ctx context.Context) IDictDatumDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dictDatumDo) ReadDB() IDictDatumDo {
	return d.Clauses(dbresolver.Read)
}

func (d dictDatumDo) WriteDB() IDictDatumDo {
	return d.Clauses(dbresolver.Write)
}

func (d dictDatumDo) Session(config *gorm.Session) IDictDatumDo {
	return d.withDO(d.DO.Session(config))
}

func (d dictDatumDo) Clauses(conds ...clause.Expression) IDictDatumDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dictDatumDo) Returning(value interface{}, columns ...string) IDictDatumDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dictDatumDo) Not(conds ...gen.Condition) IDictDatumDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dictDatumDo) Or(conds ...gen.Condition) IDictDatumDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dictDatumDo) Select(conds ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dictDatumDo) Where(conds ...gen.Condition) IDictDatumDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dictDatumDo) Order(conds ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dictDatumDo) Distinct(cols ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dictDatumDo) Omit(cols ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dictDatumDo) Join(table schema.Tabler, on ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dictDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dictDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dictDatumDo) Group(cols ...field.Expr) IDictDatumDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dictDatumDo) Having(conds ...gen.Condition) IDictDatumDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dictDatumDo) Limit(limit int) IDictDatumDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dictDatumDo) Offset(offset int) IDictDatumDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dictDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDictDatumDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dictDatumDo) Unscoped() IDictDatumDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dictDatumDo) Create(values ...*model.DictDatum) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dictDatumDo) CreateInBatches(values []*model.DictDatum, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dictDatumDo) Save(values ...*model.DictDatum) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dictDatumDo) First() (*model.DictDatum, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictDatum), nil
	}
}

func (d dictDatumDo) Take() (*model.DictDatum, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictDatum), nil
	}
}

func (d dictDatumDo) Last() (*model.DictDatum, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictDatum), nil
	}
}

func (d dictDatumDo) Find() ([]*model.DictDatum, error) {
	result, err := d.DO.Find()
	return result.([]*model.DictDatum), err
}

func (d dictDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictDatum, err error) {
	buf := make([]*model.DictDatum, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dictDatumDo) FindInBatches(result *[]*model.DictDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dictDatumDo) Attrs(attrs ...field.AssignExpr) IDictDatumDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dictDatumDo) Assign(attrs ...field.AssignExpr) IDictDatumDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dictDatumDo) Joins(fields ...field.RelationField) IDictDatumDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dictDatumDo) Preload(fields ...field.RelationField) IDictDatumDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dictDatumDo) FirstOrInit() (*model.DictDatum, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictDatum), nil
	}
}

func (d dictDatumDo) FirstOrCreate() (*model.DictDatum, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictDatum), nil
	}
}

func (d dictDatumDo) FindByPage(offset int, limit int) (result []*model.DictDatum, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dictDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dictDatumDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dictDatumDo) Delete(models ...*model.DictDatum) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dictDatumDo) withDO(do gen.Dao) *dictDatumDo {
	d.DO = *do.(*gen.DO)
	return d
}
