package request

type AddSetMultiSignReq struct {
	SpName           string   `json:"spName" form:"spName" binding:"required"`
	ChainId          int      `json:"chainId" form:"chainId" binding:"required"`
	SpToken          string   `json:"spToken" form:"spToken" binding:"required"`
	JpName           string   `json:"jpName" form:"jpName" binding:"required"`
	JpToken          string   `json:"jpToken" form:"jpToken" binding:"required"`
	SpAddress        string   `json:"spAddress" form:"spAddress" binding:"required"`
	JpAddress        string   `json:"jpAddress" form:"jpAddress" binding:"required"`
	SpHash           string   `json:"spHash" form:"spHash" binding:"required"`
	JpHash           string   `json:"jpHash" form:"jpHash" binding:"required"`
	MultiSignAccount []string `json:"multiSignAccount" form:"multiSignAccount" binding:"required"`
}
