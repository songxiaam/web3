package pledge

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建订单
func NewGetTokenInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenInfoLogic {
	return &GetTokenInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenInfoLogic) GetTokenInfo(req *types.GetTokenInfoReq) (resp *types.GetTokenInfoRes, err error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.TokenInfoModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.GetTokenInfoRes{
		Id:     one.Id,
		Symbol: one.Symbol.String,
		Logo:   one.Logo.String,
	}, nil
}
