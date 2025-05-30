package order

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建订单
func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailRes, err error) {
	// todo: add your logic here and delete this line

	return &types.GetOrderDetailRes{
		OrderId:   1,
		Count:     2,
		ProductId: 101,
	}, nil
}
