package models

type Pool struct {
	PoolID                 int      `json:"poolId"`
	SettleTime             string   `json:"settleTime"`
	EndTime                string   `json:"endTime"`
	InterestRate           string   `json:"interestRate"`
	MaxSupply              string   `json:"maxSupply"`
	LendSupply             string   `json:"lendSupply"`
	BorrowSupply           string   `json:"borrowSupply"`
	MortgageRate           string   `json:"mortgageRate"`
	LendToken              string   `json:"lendToken"`
	LendTokenSymbol        string   `json:"lendTokenSymbol"`
	BorrowToken            string   `json:"borrowToken"`
	BorrowTokenSymbol      string   `json:"borrowTokenSymbol"`
	State                  string   `json:"state"`
	SpCoin                 string   `json:"spCoin"`
	JpCoin                 string   `json:"jpCoin"`
	AutoLiquidateThreshold string   `json:"autoLiquidateThreshold"`
	PoolData               PoolData `json:"poolData"`
}
