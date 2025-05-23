package dao

import (
	"pledgev2-backend/pkg/db"
)

type PoolDateDao struct{}

func InitPoolDateDao() *PoolDateDao {
	return &PoolDateDao{}
}

func (*PoolDateDao) GetPoolDataListByChainId(chainId int) error {
	var poolDataList []TbPoolData

	err := db.MySql.Table("pool_data").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolDataList).Debug().Error
	if err != nil {
		return err
	}
	return nil
}
