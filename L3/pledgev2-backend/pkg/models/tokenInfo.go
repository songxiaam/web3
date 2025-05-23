package models

type TokenInfo struct {
	Id           int    `json:"id"`
	Symbol       string `json:"symbol"`
	Decimals     int    `json:"decimals"`
	Token        string `json:"token"`
	Logo         string `json:"logo"`
	Price        string `json:"price"`
	ChainId      string `json:"chainId"`
	abiFileExist int    `json:"abiFileExist"`
}
