package controllers

import (
	"github.com/gin-gonic/gin"
	"pledgev2-backend/pkg/common/statecode"
	"pledgev2-backend/pkg/models"
	"pledgev2-backend/pkg/models/request"
	"pledgev2-backend/pkg/models/response"
	"pledgev2-backend/pkg/service"
	"pledgev2-backend/pkg/validate"
)

type PoolBaseController struct {
}

func (c *PoolBaseController) GetPoolBaseInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.GetPoolBaseInfoReq{}
	errCode := validate.InitPoolBaseInfo().PoolBaseInfo(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}
	poolBaseList, code := service.InitPoolBaseService().GetPoolBaseList(req.ChainId)
	if code != statecode.CommonSuccess {
		res.Response(ctx, code, nil)
		return
	}
	poolBaseInfoList := models.ConvertToPoolBaseInfoList(poolBaseList)
	res.Response(ctx, statecode.CommonSuccess, poolBaseInfoList)
	return
}
