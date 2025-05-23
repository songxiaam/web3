package models

import (
	"encoding/json"
	"pledgev2-backend/pkg/dao"
)

func ConvertToMultiSign(tbMultiSign *dao.TbMultiSign) *MultiSign {
	if tbMultiSign == nil {
		return nil
	}
	return &MultiSign{
		Id:               tbMultiSign.Id,
		ChainId:          tbMultiSign.ChainId,
		JpName:           tbMultiSign.JpName,
		JpToken:          tbMultiSign.JpToken,
		JpAddress:        tbMultiSign.JpAddress,
		JpHash:           tbMultiSign.JpHash,
		SpName:           tbMultiSign.SpName,
		SpToken:          tbMultiSign.SpToken,
		SpAddress:        tbMultiSign.SpAddress,
		SpHash:           tbMultiSign.SpHash,
		MultiSignAccount: tbMultiSign.MultiSignAccount,
	}
}

func ConvertToPoolBase(tbPoolBase *dao.TbPoolBase) *PoolBase {
	if tbPoolBase == nil {
		return nil
	}
	return &PoolBase{
		PoolID:                 tbPoolBase.PoolId,
		ChainId:                tbPoolBase.ChainId,
		SettleTime:             tbPoolBase.SettleTime,
		EndTime:                tbPoolBase.EndTime,
		InterestRate:           tbPoolBase.InterestRate,
		MaxSupply:              tbPoolBase.MaxSupply,
		LendSupply:             tbPoolBase.LendSupply,
		BorrowSupply:           tbPoolBase.BorrowSupply,
		MortgageRate:           tbPoolBase.MortgageRate,
		LendToken:              tbPoolBase.LendToken,
		LendTokenInfo:          tbPoolBase.LendTokenInfo,
		BorrowToken:            tbPoolBase.BorrowToken,
		BorrowTokenInfo:        tbPoolBase.BorrowTokenInfo,
		State:                  tbPoolBase.State,
		SpCoin:                 tbPoolBase.SpCoin,
		JpCoin:                 tbPoolBase.JpCoin,
		AutoLiquidateThreshold: tbPoolBase.AutoLiquidateThreshold,
	}
}

func ConvertToPoolBaseList(tbPoolBaseList []dao.TbPoolBase) []PoolBase {
	poolBaseList := make([]PoolBase, 0, len(tbPoolBaseList))
	for _, tbPoolBase := range tbPoolBaseList {
		poolBase := ConvertToPoolBase(&tbPoolBase)
		poolBaseList = append(poolBaseList, *poolBase)
	}
	return poolBaseList
}

func ConvertToPoolData(tbPoolData *dao.TbPoolData) *PoolData {
	if tbPoolData == nil {
		return nil
	}
	return &PoolData{
		Id: tbPoolData.Id,
	}
}

func ConvertToTokenInfo(tbTokenInfo *dao.TbTokenInfo) *TokenInfo {
	if tbTokenInfo == nil {
		return nil
	}
	return &TokenInfo{
		Id:           tbTokenInfo.Id,
		Symbol:       tbTokenInfo.Symbol,
		Decimals:     tbTokenInfo.Decimals,
		Logo:         tbTokenInfo.Logo,
		Token:        tbTokenInfo.Token,
		ChainId:      tbTokenInfo.ChainId,
		abiFileExist: tbTokenInfo.AbiFileExist,
		Price:        tbTokenInfo.Price,
	}
}

func ConvertToTokenInfoList(tbTokenInfoList []dao.TbTokenInfo) []TokenInfo {
	tokenInfoList := make([]TokenInfo, 0, len(tbTokenInfoList))
	for _, tbTokenInfo := range tbTokenInfoList {
		tokenInfo := ConvertToTokenInfo(&tbTokenInfo)
		tokenInfoList = append(tokenInfoList, *tokenInfo)
	}
	return tokenInfoList
}

func ConvertToPoolBaseInfoList(poolBaseList []PoolBase) []PoolBaseInfo {
	var poolBaseInfoList = make([]PoolBaseInfo, 0)
	for _, item := range poolBaseList {
		poolBaseInfo := ConvertToPoolBaseInfo(item)
		poolBaseInfoList = append(poolBaseInfoList, poolBaseInfo)
	}
	return poolBaseInfoList
}

// 不修改值,直接用值传递
func ConvertToPoolBaseInfo(poolBase PoolBase) PoolBaseInfo {

	lendTokenInfo := LendTokenInfo{}
	err := json.Unmarshal([]byte(poolBase.LendTokenInfo), &lendTokenInfo)
	if err != nil {
	}
	borrowTokenInfo := BorrowTokenInfo{}
	err = json.Unmarshal([]byte(poolBase.BorrowTokenInfo), &borrowTokenInfo)
	if err != nil {
	}
	return PoolBaseInfo{
		PoolID:                 poolBase.PoolID,
		ChainId:                poolBase.ChainId,
		AutoLiquidateThreshold: poolBase.AutoLiquidateThreshold,
		BorrowSupply:           poolBase.BorrowSupply,
		BorrowToken:            poolBase.BorrowToken,
		BorrowTokenInfo:        borrowTokenInfo,
		EndTime:                poolBase.EndTime,
		InterestRate:           poolBase.InterestRate,
		JpCoin:                 poolBase.JpCoin,
		LendSupply:             poolBase.LendSupply,
		LendToken:              poolBase.LendToken,
		LendTokenInfo:          lendTokenInfo,
		MortgageRate:           poolBase.MortgageRate,
		MaxSupply:              poolBase.MaxSupply,
		SettleTime:             poolBase.SettleTime,
		SpCoin:                 poolBase.SpCoin,
		State:                  poolBase.State,
	}
}
