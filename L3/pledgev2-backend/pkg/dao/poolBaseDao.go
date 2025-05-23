package dao

import (
	"errors"
	"gorm.io/gorm"
	"pledgev2-backend/log"
	"pledgev2-backend/pkg/db"
	"time"
)

type PoolBaseDao struct{}

func InitPoolBaseDao() *PoolBaseDao {
	return &PoolBaseDao{}
}

func (*PoolBaseDao) Add(chainId, settleTime, endTime, interestRate, maxSupply, lendSupply, borrowSupply, mortgageRate, lendToken, lendTokenInfo,
	borrowToken, borrowTokenInfo, state, jpCoin, spCoin, autoLiquidateThreshold, lendTokenSymbol, borrowTokenSymbol string, poolId int) (int, error) {
	var tbPoolBase TbPoolBase
	now := time.Now()
	err := db.MySql.Table(tbPoolBase.TableName()).Where("chain_id_id=? and pool_id=?", chainId, poolId).First(&tbPoolBase).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tbPoolBase = TbPoolBase{
				ChainId:                chainId,
				PoolId:                 poolId,
				SettleTime:             settleTime,
				EndTime:                endTime,
				InterestRate:           interestRate,
				MaxSupply:              maxSupply,
				LendSupply:             lendSupply,
				BorrowSupply:           borrowSupply,
				MortgageRate:           mortgageRate,
				LendToken:              lendToken,
				LendTokenInfo:          lendTokenInfo,
				BorrowToken:            borrowToken,
				BorrowTokenInfo:        borrowTokenInfo,
				State:                  state,
				JpCoin:                 jpCoin,
				SpCoin:                 spCoin,
				AutoLiquidateThreshold: autoLiquidateThreshold,
				LendTokenSymbol:        lendTokenSymbol,
				BorrowTokenSymbol:      borrowTokenSymbol,
				CreatedAt:              now,
				UpdatedAt:              now,
			}
			err = db.MySql.Table(tbPoolBase.TableName()).Create(&tbPoolBase).Debug().Error
			if err != nil {
				log.Logger.Error(err.Error())
				return tbPoolBase.Id, err
			}
		}
	}
	return tbPoolBase.Id, nil
}

func (*PoolBaseDao) GetPoolBaseList(chainId int) ([]TbPoolBase, error) {
	var poolBaseList []TbPoolBase
	tbPoolBase := TbPoolBase{}
	err := db.MySql.Table(tbPoolBase.TableName()).Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolBaseList).Debug().Error
	if err != nil {
		return nil, err
	}
	return poolBaseList, nil
}

func (*PoolBaseDao) Pagination(whereCondition string) (int64, []TbPoolBase, error) {
	var poolBaseList []TbPoolBase
	var total int64
	db.MySql.Table("pool_bases").Where(whereCondition).Count(&total)

	err := db.MySql.Model(InitTbPoolBase()).Where(whereCondition).Find(&poolBaseList).Error
	if err != nil {
		return 0, nil, err
	}
	return total, poolBaseList, nil
}
