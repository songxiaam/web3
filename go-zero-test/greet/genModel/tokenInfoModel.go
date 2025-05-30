package genModel

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokenInfoModel = (*customTokenInfoModel)(nil)

type (
	// TokenInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenInfoModel.
	TokenInfoModel interface {
		tokenInfoModel
		withSession(session sqlx.Session) TokenInfoModel
	}

	customTokenInfoModel struct {
		*defaultTokenInfoModel
	}
)

// NewTokenInfoModel returns a model for the database table.
func NewTokenInfoModel(conn sqlx.SqlConn) TokenInfoModel {
	return &customTokenInfoModel{
		defaultTokenInfoModel: newTokenInfoModel(conn, nil),
	}
}

func (m *customTokenInfoModel) withSession(session sqlx.Session) TokenInfoModel {
	return NewTokenInfoModel(sqlx.NewSqlConnFromSession(session))
}

//func (m *customTokenInfoModel) FindList(ctx context.Context, startIndex, pageSize uint64) ([]TokenInfo, error) {
//	query := fmt.Sprintf("select %s from %s order by id limit ?, ?", tokenInfoRows, m.table)
//	var resp []TokenInfo
//	err := m.QueryRowCtx(ctx, &resp, query)
//	switch err {
//	case nil:
//		return resp, nil
//	case sqlx.ErrNotFound:
//		return nil, ErrNotFound
//	default:
//		return nil, err
//	}
//}

//func (m *customTokenInfoModel) TotalCount(ctx context.Context) (uint64, error) {
//	query := fmt.Sprintf("select count(*) as count from %s", m.table)
//	var count uint64
//	err := m.Que(ctx, &count, query)
//	switch err {
//	case nil:
//		return count, nil
//	case sqlx.ErrNotFound:
//		return 0, ErrNotFound
//	default:
//		return 0, err
//	}
//}

func (m *customTokenInfoModel) Search(ctx context.Context, id uint64, symbol, chainId string, startIndex, pageSize uint64) ([]TokenInfo, error) {
	query := "SELECT * FROM " + m.table + " WHERE 1=1"
	var args []interface{}
	if id != 0 {
		query += " and id = ?"
		args = append(args, id)
	}

	if symbol != "" {
		query += " and symbol = ?"
		args = append(args, symbol)
	}

	if chainId != "" {
		query += " and chain_id = ?"
		args = append(args, chainId)
	}

	if startIndex >= 0 && pageSize > 0 {
		query += " limit ?, ?"
		args = append(args, startIndex, pageSize)
	}

	var resp []TokenInfo
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
