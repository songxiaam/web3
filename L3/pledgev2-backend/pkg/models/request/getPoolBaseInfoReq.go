package request

type GetPoolBaseInfoReq struct {
	ChainId int `json:"chainId" form:"chainId" binding:"required,oneof=56 97"`
}
