package models

type PoolBaseInfo struct {
	PoolID                 int             `json:"poolId"`
	ChainId                string          `json:"chainId"`
	AutoLiquidateThreshold string          `json:"auto_liquidate_threshold"`
	BorrowSupply           string          `json:"borrow_supply"`
	BorrowToken            string          `json:"borrow_token"`
	BorrowTokenInfo        BorrowTokenInfo `json:"borrow_token_info"`
	EndTime                string          `json:"end_time"`
	InterestRate           string          `json:"interest_rate"`
	JpCoin                 string          `json:"jp_coin"`
	LendSupply             string          `json:"lend_supply"`
	LendToken              string          `json:"lend_token"`
	LendTokenInfo          LendTokenInfo   `json:"lend_token_info"`
	MortgageRate           string          `json:"mortgage_rate"`
	MaxSupply              string          `json:"max_supply"`
	SettleTime             string          `json:"settle_time"`
	SpCoin                 string          `json:"sp_coin"`
	State                  string          `json:"state"`
}

type PoolBase struct {
	PoolID                 int    `json:"poolId"`
	ChainId                string `json:"chainId"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold"`
	BorrowSupply           string `json:"borrowSupply"`
	BorrowToken            string `json:"borrowToken"`
	BorrowTokenInfo        string `json:"borrowTokenInfo"`
	EndTime                string `json:"endTime"`
	InterestRate           string `json:"interestRate"`
	JpCoin                 string `json:"jpCoin"`
	LendSupply             string `json:"lendSupply"`
	LendToken              string `json:"lendToken"`
	LendTokenInfo          string `json:"lendTokenInfo"`
	MortgageRate           string `json:"mortgageRate"`
	MaxSupply              string `json:"maxSupply"`
	SettleTime             string `json:"settleTime"`
	SpCoin                 string `json:"spCoin"`
	State                  string `json:"state"`
}

type BorrowTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}

type LendTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`
	TokenLogo  string `json:"tokenLogo"`
	TokenName  string `json:"tokenName"`
	TokenPrice string `json:"tokenPrice"`
}
