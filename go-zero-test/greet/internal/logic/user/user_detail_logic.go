package user

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户
func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.GetUserDetailReq) (resp *types.GetUserDetailRes, err error) {
	// todo: add your logic here and delete this line
	return &types.GetUserDetailRes{
		UserId: 201,
		Name:   "MikeZhao",
	}, nil
}
