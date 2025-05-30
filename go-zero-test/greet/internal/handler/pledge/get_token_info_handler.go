package pledge

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/internal/logic/pledge"
	"greet/internal/svc"
	"greet/internal/types"
)

// 创建订单
func GetTokenInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pledge.NewGetTokenInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetTokenInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
