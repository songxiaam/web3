// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type GetTokenInfoListRes struct {
	List  []TokenInfo `json:"list"`
	Total uint64      `json:"total"`
}

type PoolBase struct {
	Id                int64  `json:"id"`
	ChainId           string `json:"chainId"`
	LendTokenSymbol   string `json:"lendTokenSymbol"`
	BorrowTokenSymbol string `json:"borrowTokenSymbol"`
}

type SearchRes struct {
	List  []TokenInfo `json:"list"`
	Total uint64      `json:"total"`
}

type TokenInfo struct {
	Id         uint64 `json:"id"`
	Symbol     string `json:"symbol"`
	Logo       string `json:"logo"`
	Token      string `json:"token"`
	ChainId    string `json:"chain_id"`
	CustomCode uint64 `json:"customCode"`
}
