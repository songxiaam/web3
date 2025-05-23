package response

import (
	"github.com/gin-gonic/gin"
	"pledgev2-backend/pkg/common/statecode"
)

type Gin struct {
	Res *gin.Context
}

type Page struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// 分页数据
func (g *Gin) ResponsePages(c *gin.Context, code int, totalCount int, data interface{}) {
	lang := statecode.LangZh
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}

	rsp := Page{
		Code:  code, //业务状态代码
		Msg:   statecode.GetMsg(code, lang),
		Total: totalCount,
		Data:  data,
	}
	g.Res.JSON(statecode.HttpStatusSuccess, rsp)
}

func (g *Gin) Response(c *gin.Context, code int, data interface{}, httpStatus ...int) {
	lang := statecode.LangEn
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	rsp := Response{
		Code: code,
		Msg:  statecode.GetMsg(code, lang),
		Data: data,
	}
	HttpStatus := statecode.HttpStatusSuccess
	if len(httpStatus) > 0 {
		HttpStatus = httpStatus[0]
	}
	g.Res.JSON(HttpStatus, rsp)
	return
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
