package request

type Search struct {
	ChainID         int    `json:"chainID" form:"chainID" binding:"required"`
	LendTokenSymbol string `json:"lend_token_symbol" form:"lend_token_symbol" binding:"required"`
	State           string `json:"state" form:"state" binding:"required"`
	Page            int    `json:"page" form:"page" binding:"required"`
	PageSize        int    `json:"pageSize" form:"pageSize" binding:"required"`
}
