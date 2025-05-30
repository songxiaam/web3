package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/internal/logic/order"
	"greet/internal/svc"
	"greet/internal/types"
)

// 创建订单
func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
