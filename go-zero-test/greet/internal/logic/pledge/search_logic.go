package pledge

import (
	"context"
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

func (l *SearchLogic) Search(req *types.SearchReq) (resp *response.BaseResp[*types.SearchRes], err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.TokenInfoModel.Search(l.ctx, req.Id, req.Symbol, req.ChainId, req.StartIndex, req.PageSize)
	if err != nil {
		return nil, err
	}
	count, err := l.svcCtx.TokenInfoModel.TotalCount(l.ctx)

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
