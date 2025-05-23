package models

type PoolData struct {
	Id                    int    `json:"id"`
	PoolId                string `json:"poolId"`
	ChainId               string `json:"chainId"`
	SettleAmountLend      string `json:"settleAmountLend"`
	SettleAmountBorrow    string `json:"settleAmountBorrow"`
	FinishAmountLend      string `json:"finishAMountLend"`
	FinishAmountBorrow    string `json:"finishAMountBorrow"`
	LiquidateAmountLend   string `json:"liquidateAMountLend"`
	LiquidateAmountBorrow string `json:"liquidateAMountBorrow"`
}

type PoolDataInfoRes struct {
	Index    int      `json:"index"`
	PoolData PoolData `json:"poolData"`
}
