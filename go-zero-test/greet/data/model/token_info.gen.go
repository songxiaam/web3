// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

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

func newTokenInfo(db *gorm.DB, opts ...gen.DOOption) tokenInfo {
	_tokenInfo := tokenInfo{}

	_tokenInfo.tokenInfoDo.UseDB(db, opts...)
	_tokenInfo.tokenInfoDo.UseModel(&model.TokenInfo{})

	tableName := _tokenInfo.tokenInfoDo.TableName()
	_tokenInfo.ALL = field.NewAsterisk(tableName)
	_tokenInfo.ID = field.NewInt32(tableName, "id")
	_tokenInfo.Symbol = field.NewString(tableName, "symbol")
	_tokenInfo.Logo = field.NewString(tableName, "logo")
	_tokenInfo.Price = field.NewString(tableName, "price")
	_tokenInfo.Token = field.NewString(tableName, "token")
	_tokenInfo.ChainID = field.NewString(tableName, "chain_id")
	_tokenInfo.AbiFileExist = field.NewInt32(tableName, "abi_file_exist")
	_tokenInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_tokenInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tokenInfo.Decimals = field.NewInt32(tableName, "decimals")
	_tokenInfo.Desc = field.NewString(tableName, "desc")

	_tokenInfo.fillFieldMap()

	return _tokenInfo
}

type tokenInfo struct {
	tokenInfoDo

	ALL          field.Asterisk
	ID           field.Int32
	Symbol       field.String
	Logo         field.String
	Price        field.String
	Token        field.String
	ChainID      field.String
	AbiFileExist field.Int32
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Decimals     field.Int32
	Desc         field.String

	fieldMap map[string]field.Expr
}

func (t tokenInfo) Table(newTableName string) *tokenInfo {
	t.tokenInfoDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tokenInfo) As(alias string) *tokenInfo {
	t.tokenInfoDo.DO = *(t.tokenInfoDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tokenInfo) updateTableName(table string) *tokenInfo {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Symbol = field.NewString(table, "symbol")
	t.Logo = field.NewString(table, "logo")
	t.Price = field.NewString(table, "price")
	t.Token = field.NewString(table, "token")
	t.ChainID = field.NewString(table, "chain_id")
	t.AbiFileExist = field.NewInt32(table, "abi_file_exist")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.Decimals = field.NewInt32(table, "decimals")
	t.Desc = field.NewString(table, "desc")

	t.fillFieldMap()

	return t
}

func (t *tokenInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tokenInfo) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["symbol"] = t.Symbol
	t.fieldMap["logo"] = t.Logo
	t.fieldMap["price"] = t.Price
	t.fieldMap["token"] = t.Token
	t.fieldMap["chain_id"] = t.ChainID
	t.fieldMap["abi_file_exist"] = t.AbiFileExist
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["decimals"] = t.Decimals
	t.fieldMap["desc"] = t.Desc
}

func (t tokenInfo) clone(db *gorm.DB) tokenInfo {
	t.tokenInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tokenInfo) replaceDB(db *gorm.DB) tokenInfo {
	t.tokenInfoDo.ReplaceDB(db)
	return t
}

type tokenInfoDo struct{ gen.DO }

type ITokenInfoDo interface {
	gen.SubQuery
	Debug() ITokenInfoDo
	WithContext(ctx context.Context) ITokenInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITokenInfoDo
	WriteDB() ITokenInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITokenInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITokenInfoDo
	Not(conds ...gen.Condition) ITokenInfoDo
	Or(conds ...gen.Condition) ITokenInfoDo
	Select(conds ...field.Expr) ITokenInfoDo
	Where(conds ...gen.Condition) ITokenInfoDo
	Order(conds ...field.Expr) ITokenInfoDo
	Distinct(cols ...field.Expr) ITokenInfoDo
	Omit(cols ...field.Expr) ITokenInfoDo
	Join(table schema.Tabler, on ...field.Expr) ITokenInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITokenInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITokenInfoDo
	Group(cols ...field.Expr) ITokenInfoDo
	Having(conds ...gen.Condition) ITokenInfoDo
	Limit(limit int) ITokenInfoDo
	Offset(offset int) ITokenInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITokenInfoDo
	Unscoped() ITokenInfoDo
	Create(values ...*model.TokenInfo) error
	CreateInBatches(values []*model.TokenInfo, batchSize int) error
	Save(values ...*model.TokenInfo) error
	First() (*model.TokenInfo, error)
	Take() (*model.TokenInfo, error)
	Last() (*model.TokenInfo, error)
	Find() ([]*model.TokenInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TokenInfo, err error)
	FindInBatches(result *[]*model.TokenInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TokenInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITokenInfoDo
	Assign(attrs ...field.AssignExpr) ITokenInfoDo
	Joins(fields ...field.RelationField) ITokenInfoDo
	Preload(fields ...field.RelationField) ITokenInfoDo
	FirstOrInit() (*model.TokenInfo, error)
	FirstOrCreate() (*model.TokenInfo, error)
	FindByPage(offset int, limit int) (result []*model.TokenInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITokenInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tokenInfoDo) Debug() ITokenInfoDo {
	return t.withDO(t.DO.Debug())
}

func (t tokenInfoDo) WithContext(ctx context.Context) ITokenInfoDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tokenInfoDo) ReadDB() ITokenInfoDo {
	return t.Clauses(dbresolver.Read)
}

func (t tokenInfoDo) WriteDB() ITokenInfoDo {
	return t.Clauses(dbresolver.Write)
}

func (t tokenInfoDo) Session(config *gorm.Session) ITokenInfoDo {
	return t.withDO(t.DO.Session(config))
}

func (t tokenInfoDo) Clauses(conds ...clause.Expression) ITokenInfoDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tokenInfoDo) Returning(value interface{}, columns ...string) ITokenInfoDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tokenInfoDo) Not(conds ...gen.Condition) ITokenInfoDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tokenInfoDo) Or(conds ...gen.Condition) ITokenInfoDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tokenInfoDo) Select(conds ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tokenInfoDo) Where(conds ...gen.Condition) ITokenInfoDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tokenInfoDo) Order(conds ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tokenInfoDo) Distinct(cols ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tokenInfoDo) Omit(cols ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tokenInfoDo) Join(table schema.Tabler, on ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tokenInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tokenInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tokenInfoDo) Group(cols ...field.Expr) ITokenInfoDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tokenInfoDo) Having(conds ...gen.Condition) ITokenInfoDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tokenInfoDo) Limit(limit int) ITokenInfoDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tokenInfoDo) Offset(offset int) ITokenInfoDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tokenInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITokenInfoDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tokenInfoDo) Unscoped() ITokenInfoDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tokenInfoDo) Create(values ...*model.TokenInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tokenInfoDo) CreateInBatches(values []*model.TokenInfo, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tokenInfoDo) Save(values ...*model.TokenInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tokenInfoDo) First() (*model.TokenInfo, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TokenInfo), nil
	}
}

func (t tokenInfoDo) Take() (*model.TokenInfo, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TokenInfo), nil
	}
}

func (t tokenInfoDo) Last() (*model.TokenInfo, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TokenInfo), nil
	}
}

func (t tokenInfoDo) Find() ([]*model.TokenInfo, error) {
	result, err := t.DO.Find()
	return result.([]*model.TokenInfo), err
}

func (t tokenInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TokenInfo, err error) {
	buf := make([]*model.TokenInfo, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tokenInfoDo) FindInBatches(result *[]*model.TokenInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tokenInfoDo) Attrs(attrs ...field.AssignExpr) ITokenInfoDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tokenInfoDo) Assign(attrs ...field.AssignExpr) ITokenInfoDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tokenInfoDo) Joins(fields ...field.RelationField) ITokenInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tokenInfoDo) Preload(fields ...field.RelationField) ITokenInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tokenInfoDo) FirstOrInit() (*model.TokenInfo, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TokenInfo), nil
	}
}

func (t tokenInfoDo) FirstOrCreate() (*model.TokenInfo, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TokenInfo), nil
	}
}

func (t tokenInfoDo) FindByPage(offset int, limit int) (result []*model.TokenInfo, count int64, err error) {
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

func (t tokenInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tokenInfoDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tokenInfoDo) Delete(models ...*model.TokenInfo) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tokenInfoDo) withDO(do gen.Dao) *tokenInfoDo {
	t.DO = *do.(*gen.DO)
	return t
}
