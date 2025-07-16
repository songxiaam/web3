package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"smart-route/pkg/data/entity"
)

// CreateToken 创建新 Token
func CreateToken(ctx context.Context, db *gorm.DB, token *entity.Token) error {
	if token.ID == uuid.Nil {
		token.ID = uuid.New()
	}
	return db.WithContext(ctx).Create(token).Error
}

// ListTokens 查询 token 列表，可根据 chain, chainId, name, symbol 筛选
func ListTokens(
	ctx context.Context,
	db *gorm.DB,
	chain string,
	chainId *int,
	name, symbol string,
	page, pageSize int,
) ([]entity.Token, int64, error) {
	var tokens []entity.Token
	var total int64

	query := db.WithContext(ctx).Model(&entity.Token{})

	if chain != "" {
		query = query.Where("chain = ?", chain)
	}
	if chainId != nil {
		query = query.Where("chain_id = ?", *chainId)
	}
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if symbol != "" {
		query = query.Where("symbol = ?", symbol)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Find(&tokens).Error; err != nil {
		return nil, 0, err
	}
	return tokens, total, nil
}
