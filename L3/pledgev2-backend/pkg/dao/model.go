package dao

import (
	"time"
)

type TbMultiSign struct {
	Id               int32     `gorm:"column:id;primaryKey"`
	SpName           string    `json:"spName" gorm:"column:sp_name"`
	ChainId          int       `json:"chainId" gorm:"column:chain_id"`
	SpToken          string    `json:"spToken" gorm:"column:sp_token"`
	JpName           string    `json:"jpName" gorm:"column:jp_name"`
	JpToken          string    `json:"jpToken" gorm:"column:jp_token"`
	SpAddress        string    `json:"spAddress" gorm:"column:sp_address"`
	JpAddress        string    `json:"jpAddress" gorm:"column:jp_address"`
	SpHash           string    `json:"spHash" gorm:"column:sp_hash"`
	JpHash           string    `json:"jpHash" gorm:"column:jp_hash"`
	MultiSignAccount string    `json:"multiSignAccount" gorm:"column:multi_sign_account"`
	CreateAt         time.Time `json:"createAt" gorm:"column:create_at"`
	UpdateAt         time.Time `json:"updateAt" gorm:"column:update_at"`
}

func InitTbMultiSign() *TbMultiSign {
	return &TbMultiSign{}
}

func (*TbMultiSign) TableName() string {
	return "multi_sign"
}

type TbPoolBase struct {
	Id                     int       `json:"id" gorm:"column:id;primaryKey"`
	PoolId                 int       `json:"poolId" gorm:"column:pool_id"`
	ChainId                string    `json:"chainId" gorm:"column:chain_id"`
	SettleTime             string    `json:"settleTime" gorm:"column:settle_time"`
	EndTime                string    `json:"endTime" gorm:"column:end_time"`
	InterestRate           string    `json:"interestRate" gorm:"column:interest_rate"`
	MaxSupply              string    `json:"maxSupply" gorm:"max_supply:"`
	LendSupply             string    `json:"lendSupply" gorm:"column:lend_supply"`
	BorrowSupply           string    `json:"borrowSupply" gorm:"column:borrow_supply"`
	MortgageRate           string    `json:"mortgageRate" gorm:"column:martgage_rate"`
	LendToken              string    `json:"lendToken" gorm:"column:lend_token"`
	LendTokenInfo          string    `json:"lendTokenInfo" gorm:"column:lend_token_info"`
	BorrowToken            string    `json:"borrowToken" gorm:"column:borrow_token"`
	BorrowTokenInfo        string    `json:"borrowToken_Info" gorm:"column:borrow_token_info"`
	State                  string    `json:"state" gorm:"column:state"`
	SpCoin                 string    `json:"spCoin" gorm:"column:sp_coin"`
	JpCoin                 string    `json:"jpCoin" gorm:"column:jp_coin"`
	LendTokenSymbol        string    `json:"lendTokenSymbol" gorm:"column:lend_token_symbol"`
	BorrowTokenSymbol      string    `json:"borrowTokenSymbol" gorm:"column:borrow_token_symbol"`
	AutoLiquidateThreshold string    `json:"autoLiquidateThreshold" gorm:"column:auto_liquidate_threshold"`
	CreatedAt              time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt              time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func InitTbPoolBase() *TbPoolBase {
	return &TbPoolBase{}
}

func (*TbPoolBase) TableName() string {
	return "pool_base"
}

type TbPoolData struct {
	Id                      int       `json:"id" gorm:"column:id;primaryKey"`
	PoolId                  string    `json:"poolId" gorm:"column:pool_id"`
	ChainId                 string    `json:"chainId" gorm:"column:chain_id"`
	FinishAmountBorrow      string    `json:"finishAmountBorrow" gorm:"column:finish_amount_borrow"`
	FinishAmountLend        string    `json:"finishAmountLend" gorm:"column:finish_amount_lend"`
	LiquidationAmountBorrow string    `json:"liquidationAmountBorrow" gorm:"column:liquidation_amount_borrow"`
	LiquidationAmountLend   string    `json:"liquidationAmountLend" gorm:"column:liquidation_amount_lend"`
	SettleAmountBorrow      string    `json:"settleAmountBorrow" gorm:"column:settle_amount_borrow"`
	SettleAmountLend        string    `json:"settleAmountLend" gorm:"column:settle_amount_lend"`
	CreatedAt               time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt               time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (*TbPoolData) TableName() string {
	return "pool_data"
}

type TbTokenInfo struct {
	Id           int       `json:"id" gorm:"column:id;primaryKey"`
	Logo         string    `json:"logo" gorm:"column:logo"`
	Token        string    `json:"token" gorm:"column:token"`
	Symbol       string    `json:"symbol" gorm:"column:symbol"`
	ChainId      string    `json:"chainId" gorm:"column:chain_id"`
	Price        string    `json:"price" gorm:"column:price"`
	Decimals     int       `json:"decimals" gorm:"column:decimals"`
	AbiFileExist int       `json:"abiFileExist" gorm:"column:abi_file_exist"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (*TbTokenInfo) TableName() string {
	return "token_info"
}

func InitTbTokenInfo() *TbTokenInfo {
	return &TbTokenInfo{}
}
