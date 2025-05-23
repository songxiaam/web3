package dao

import (
	"errors"
	"gorm.io/gorm"
	"pledgev2-backend/pkg/db"
	"time"
)

type TokenInfoDao struct{}

func InitTokenInfoDao() *TokenInfoDao {
	return &TokenInfoDao{}
}

func (*TokenInfoDao) Add(logo, token, symbol, chainId, price string, decimals, abiFileExist int) (int, error) {
	tbTokenInfo := InitTbTokenInfo()
	createTime := time.Now()

	// 查询是否已存在
	err := db.MySql.Table(tbTokenInfo.TableName()).Where("chain_id=? and token=?", chainId, token).Debug().First(tbTokenInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tbTokenInfo = &TbTokenInfo{
				Logo:         logo,
				Token:        token,
				Symbol:       symbol,
				ChainId:      chainId,
				Price:        price,
				Decimals:     decimals,
				AbiFileExist: abiFileExist,
				CreatedAt:    createTime,
				UpdatedAt:    createTime,
			}
			db.MySql.Table(tbTokenInfo.TableName()).Create(tbTokenInfo)
			return tbTokenInfo.Id, nil
		} else {
			return -1, errors.New("token_info record select err " + err.Error())
		}
	}
	return -1, err
}

func (*TokenInfoDao) GetTokenInfoList(chainId, token string, startIndex, pageSize int) ([]TbTokenInfo, error) {
	var tokenInfoInfo []TbTokenInfo

	dbQuery := db.MySql.Model(&TbTokenInfo{})
	if chainId != "" {
		dbQuery.Where("chain_id=?", chainId)
	}
	if token != "" {
		dbQuery.Where("token=?", token)
	}

	// 分页
	//offset := (page - 1) * pageSize
	err := dbQuery.Limit(pageSize).Offset(startIndex).Find(&tokenInfoInfo).Debug().Error
	if err != nil {
		return nil, errors.New("record select err " + err.Error())
	}
	return tokenInfoInfo, nil
}
