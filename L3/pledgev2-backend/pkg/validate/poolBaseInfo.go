package validate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledgev2-backend/pkg/common/statecode"
	"pledgev2-backend/pkg/models/request"
)

type PoolBaseInfo struct{}

func InitPoolBaseInfo() *PoolBaseInfo {
	return &PoolBaseInfo{}
}

func (p *PoolBaseInfo) PoolBaseInfo(c *gin.Context, req *request.GetPoolBaseInfoReq) int {
	// 请求中绑定参数到req(解析请求, 获取请求req)
	err := c.ShouldBind(&req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, e := range errs {
			if e.Field() == "ChainId" && e.Tag() == "required" {
				return statecode.ChainIdErr
			}
		}
		return statecode.ParameterEmptyErr
	}

	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.ChainIdErr
	}
	return statecode.CommonSuccess
}
