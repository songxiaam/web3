package model

// Token 资产业务模型结构体
type Token struct {
	ID        string `json:"id"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Decimals  int    `json:"decimals"`
	Chain     string `json:"chain"`
	ChainId   int64  `json:"chainId"`
	IsStable  bool   `json:"isStable"`
	IsNative  bool   `json:"isNative"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
