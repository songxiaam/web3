package pledge

import (
	"context"
	"greet/internal/model"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建订单
func NewGetTokenInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenInfoListLogic {
	return &GetTokenInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenInfoListLogic) GetTokenInfoList(req *types.GetTokenInfoListReq) (*types.BaseResp, error) {
	// todo: add your logic here and delete this line
	logx.Info("GetTokenInfoList")
	logx.Error("logx.error")
	logx.Info("req: %v", req)
	var list []model.TokenInfo
	err := l.svcCtx.TokenInfoModel.FindList(l.ctx, req.StartIndex, req.PageSize, &list)
	if err != nil {
		return nil, err
	}
	var count uint64
	err = l.svcCtx.TokenInfoModel.TotalCount(l.ctx, 0, "", "", &count)
	if err != nil {
		return nil, err
	}
	respList := make([]types.TokenInfo, 0, len(list))
	for i, m := range list {
		tokenInfo := types.TokenInfo{
			Id:         m.Id,
			Symbol:     m.Symbol.String,
			Logo:       m.Logo.String,
			Token:      m.Token.String,
			ChainId:    m.ChainId,
			CustomCode: uint64(i),
		}
		respList = append(respList, tokenInfo)
	}

	result := types.BaseResp{
		Code: 0,
		Msg:  "success",
		Data: types.GetTokenInfoListRes{
			List:  respList,
			Total: count,
		},
	}
	return &types.BaseResp{
		Code: result.Code,
		Msg:  result.Msg,
		Data: result.Data,
	}, nil
}
