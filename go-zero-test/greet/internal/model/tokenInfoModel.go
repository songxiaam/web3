package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TokenInfoModel = (*customTokenInfoModel)(nil)

type (
	// TokenInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTokenInfoModel.
	TokenInfoModel interface {
		tokenInfoModel
		Search(ctx context.Context, id uint64, symbol, chainId string, startIndex, pageSize uint64, resp []TokenInfo) error
		FindList(ctx context.Context, startIndex, pageSize uint64, resp *[]TokenInfo) error
		TotalCount(ctx context.Context, id uint64, symbol, chainId string, count *uint64) error
	}

	customTokenInfoModel struct {
		*defaultTokenInfoModel
	}
)

// NewTokenInfoModel returns a model for the database table.
func NewTokenInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TokenInfoModel {
	return &customTokenInfoModel{
		defaultTokenInfoModel: newTokenInfoModel(conn, c, opts...),
	}
}

func (m *customTokenInfoModel) Search(ctx context.Context, id uint64, symbol, chainId string, startIndex, pageSize uint64, resp []TokenInfo) error {
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

	//var resp []TokenInfo
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, sqlx.ErrNotFound):
		return ErrNotFound
	default:
		return nil
	}
}

func (m *customTokenInfoModel) FindList(ctx context.Context, startIndex, pageSize uint64, resp *[]TokenInfo) error {
	query := "SELECT * FROM " + m.table + " order by id desc limit ?,?"
	err := m.QueryRowsNoCacheCtx(ctx, resp, query, startIndex, pageSize)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (m *customTokenInfoModel) TotalCount(ctx context.Context, id uint64, symbol, chainId string, count *uint64) error {
	query := "SELECT COUNT(*) FROM " + m.table + " WHERE 1=1"
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
	err := m.QueryRowNoCacheCtx(ctx, count, query, args...)
	if err != nil {
		return err
	}
	return nil
}
