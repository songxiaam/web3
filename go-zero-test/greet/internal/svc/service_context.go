package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"greet/internal/config"
	"greet/internal/middleware"
	"greet/internal/model"
)

type ServiceContext struct {
	Config            config.Config
	PoolBaseModel     model.PoolBaseModel
	TokenInfoModel    model.TokenInfoModel
	TestMiddleware    rest.Middleware
	RequestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		TestMiddleware:    middleware.NewTestMiddleware().Handle,
		RequestMiddleware: middleware.NewTestMiddleware().Handle,
		PoolBaseModel:     model.NewPoolBaseModel(sqlx.NewMysql(c.DB.DataSource)),
		TokenInfoModel:    model.NewTokenInfoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
