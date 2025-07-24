package service

import (
	"context"
	"gorm.io/gorm"
	"smart-route/internal/model"
	"smart-route/pkg/convert"
	"smart-route/pkg/data/entity"
	"smart-route/pkg/data/repository"
)

type TokenService struct {
	db *gorm.DB
}

func NewTokenService(db *gorm.DB) *TokenService {
	return &TokenService{db: db}
}

// ListTokens 查询 token 列表，可根据 chain, chainId, name, symbol 筛选
func (s *TokenService) ListTokens(
	ctx context.Context,
	chain string,
	chainId *int,
	name, symbol string,
	page, pageSize int,
) ([]model.Token, int64, error) {
	tokenEntities, total, err := repository.ListTokens(ctx, s.db, chain, chainId, name, symbol, page, pageSize)
	var tokens []model.Token
	convert.CopyStructFieldsSlice(&tokens, tokenEntities)
	return tokens, total, err
}

func (s *TokenService) CreateToken(ctx context.Context, token *model.Token) error {
	var tokenEntity *entity.Token
	convert.CopyStructFields(tokenEntity, token)
	return repository.CreateToken(ctx, s.db, tokenEntity)
}
