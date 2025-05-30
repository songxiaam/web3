package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"greet/internal/config"
	"greet/internal/model"
)

type ServiceContext struct {
	Config         config.Config
	PoolBaseModel  model.PoolBaseModel
	TokenInfoModel model.TokenInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		PoolBaseModel:  model.NewPoolBaseModel(sqlx.NewMysql(c.DB.DataSource)),
		TokenInfoModel: model.NewTokenInfoModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
