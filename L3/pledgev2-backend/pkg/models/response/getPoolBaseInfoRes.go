package response

import "pledgev2-backend/pkg/models"

type GetPoolBaseInfoRes struct {
	PoolBaseInfo []models.PoolBaseInfo
}

type PoolBaseInfoRes struct {
	Index    int                 `json:"index"`
	PoolData models.PoolBaseInfo `json:"poolData"`
}
