package api

import "smart-route/internal/model"

// TokenListRequest Token 列表请求结构体
type TokenListRequest struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
	Chain    string `form:"chain"`
	Symbol   string `form:"symbol"`
	Name     string `form:"name"`
	ChainId  int    `form:"chainId"`
}

// TokenListResponse Token 列表响应结构体
type TokenListResponse struct {
	Tokens   []model.Token `json:"tokens"`
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}

// TokenItem Token 列表项
type TokenItem struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Chain    string `json:"chain"`
	ChainId  int64  `json:"chain_id"`
	IsStable bool   `json:"is_stable"`
	IsNative bool   `json:"is_native"`
}

// TokenCreateRequest 资产录入请求
type TokenCreateRequest struct {
	Symbol   string `json:"symbol" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals" binding:"required"`
	Chain    string `json:"chain" binding:"required"`
	ChainId  int64  `json:"chain_id" binding:"required"`
	IsStable bool   `json:"is_stable"`
	IsNative bool   `json:"is_native"`
}

// TokenCreateResponse 资产录入响应
type TokenCreateResponse struct {
	ID string `json:"id"`
}
