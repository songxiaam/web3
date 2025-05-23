package request

type TokenList struct {
	ChainId int `json:"chainId" form:"chainId" binding:"required"`
}
