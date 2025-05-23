package service

import (
	"go.uber.org/zap"
	"pledgev2-backend/log"
	"pledgev2-backend/pkg/common/statecode"
	"pledgev2-backend/pkg/dao"
	"pledgev2-backend/pkg/models"
)

type poolBaseService struct{}

func InitPoolBaseService() *poolBaseService {
	return &poolBaseService{}
}

func (p *poolBaseService) GetPoolBaseList(chainId int) ([]models.PoolBase, int) {
	//err := models.NewPoolBases().PoolBaseInfo(chainId, result)

	tbPoolBaseList, err := dao.InitPoolBaseDao().GetPoolBaseList(chainId)
	if err != nil {
		log.Logger.Error("poolService.PoolBaseInfo err", zap.Error(err))
		return nil, statecode.CommonErrServerErr
	}
	poolBaseList := models.ConvertToPoolBaseList(tbPoolBaseList)
	return poolBaseList, statecode.CommonSuccess
}
