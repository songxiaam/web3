package poolBase

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/internal/logic/poolBase"
	"greet/internal/svc"
	"greet/internal/types"
)

// 获取PoolBase
func GetPoolBaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPoolBaseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := poolBase.NewGetPoolBaseLogic(r.Context(), svcCtx)
		resp, err := l.GetPoolBase(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
