package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"pledgev2-backend/pkg/db"
)

type MultiSignDao struct{}

// 相当于init
func InitMultiSignDao() *MultiSignDao {
	return &MultiSignDao{}
}

// 相当于实例方法
func (*MultiSignDao) Add(chainId int, spName, spToken, jpName, jpToken, spAddress, jpAddress, spHash, jpHash, multiSignAccount string) error {
	err := db.MySql.Table(InitTbMultiSign().TableName()).Where("chain_id=?", chainId).Delete(&TbMultiSign{}).Debug().Error
	if err != nil {
		return fmt.Errorf("failed to delete old records: %w", err)
	}
	err = db.MySql.Table(InitTbMultiSign().TableName()).Create(&TbMultiSign{
		ChainId:          chainId,
		SpName:           spName,
		SpToken:          spToken,
		JpName:           jpName,
		JpToken:          jpToken,
		SpAddress:        spAddress,
		JpAddress:        jpAddress,
		SpHash:           spHash,
		JpHash:           jpHash,
		MultiSignAccount: multiSignAccount,
	}).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (*MultiSignDao) Get(chainId int) (*TbMultiSign, error) {
	tbMultiSign := &TbMultiSign{}
	err := db.MySql.Table("multi_sign").Where("chain_id=?", chainId).First(&tbMultiSign).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else {
			return nil, errors.New("record select err " + err.Error())
		}
	}
	return tbMultiSign, nil
}
