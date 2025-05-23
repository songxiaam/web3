package request

type PoolDataInfo struct {
	ChainId int `json:"chainId" form:"chainId" binding:"required"`
}
