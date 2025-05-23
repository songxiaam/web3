package request

type GetMultiSign struct {
	ChainId int `json:"chainId" form:"chainId" binding:"required"`
}
