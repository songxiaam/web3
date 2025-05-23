package models

// gorm 映射数据库
type MultiSign struct {
	Id               int32  `json:"id"`
	SpName           string `json:"spName"`
	ChainId          int    `json:"chainId"`
	SpToken          string `json:"spToken"`
	JpName           string `json:"jpName"`
	JpToken          string `json:"jpToken"`
	SpAddress        string `json:"spAddress"`
	JpAddress        string `json:"jpAddress"`
	SpHash           string `json:"spHash"`
	JpHash           string `json:"jpHash"`
	MultiSignAccount string `json:"multiSignAccount"`
}
