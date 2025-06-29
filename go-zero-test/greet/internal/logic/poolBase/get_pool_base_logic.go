package poolBase

import (
	"context"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolBaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取PoolBase
func NewGetPoolBaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolBaseLogic {
	return &GetPoolBaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolBaseLogic) GetPoolBase(req *types.GetPoolBaseReq) (resp *types.GetPoolBaseRes, err error) {
	// todo: add your logic here and delete this line

	poolBase, err := l.svcCtx.PoolBaseModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, err
	}
	return &types.GetPoolBaseRes{
		PoolBase: types.PoolBase{
			Id:                poolBase.Id,
			ChainId:           poolBase.ChainId,
			LendTokenSymbol:   poolBase.LendTokenSymbol.String,
			BorrowTokenSymbol: poolBase.BorrowTokenSymbol.String,
		},
	}, nil
}
