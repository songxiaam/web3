package pledge

import (
	"context"
	"greet/internal/model"
	"greet/internal/pkg/response"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取TokenList
func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*response.BaseResp[*types.SearchRes], error) {
	// todo: add your logic here and delete this line
	var list []model.TokenInfo
	err := l.svcCtx.TokenInfoModel.Search(l.ctx, req.Id, req.Symbol, req.ChainId, req.StartIndex, req.PageSize, list)
	if err != nil {
		return nil, err
	}

	var count uint64
	err = l.svcCtx.TokenInfoModel.TotalCount(l.ctx, req.Id, req.Symbol, req.ChainId, &count)

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

	result := types.SearchRes{
		List:  respList,
		Total: count,
	}
	return &response.BaseResp[*types.SearchRes]{
		Code: 0,
		Msg:  "success",
		Data: &result,
	}, nil
}
