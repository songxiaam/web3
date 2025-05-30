package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PoolBaseModel = (*customPoolBaseModel)(nil)

type (
	// PoolBaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPoolBaseModel.
	PoolBaseModel interface {
		poolBaseModel
		withSession(session sqlx.Session) PoolBaseModel
	}

	customPoolBaseModel struct {
		*defaultPoolBaseModel
	}
)

// NewPoolBaseModel returns a model for the database table.
func NewPoolBaseModel(conn sqlx.SqlConn) PoolBaseModel {
	return &customPoolBaseModel{
		defaultPoolBaseModel: newPoolBaseModel(conn),
	}
}

func (m *customPoolBaseModel) withSession(session sqlx.Session) PoolBaseModel {
	return NewPoolBaseModel(sqlx.NewSqlConnFromSession(session))
}
