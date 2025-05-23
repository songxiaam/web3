package service

type SearchService struct{}

func NewSearchService() *SearchService {
	return &SearchService{}
}

//func (p *SearchService) Search(request *request.Search) (int, int64, []models.Pool) {
//	whereCondition := fmt.Sprintf("chain_id=?", request.ChainID)
//	if request.LendTokenSymbol != "" {
//		whereCondition += fmt.Sprintf(` and lend_token_symbol='%v'`, request.LendTokenSymbol)
//	}
//	if request.State != "" {
//		whereCondition += fmt.Sprintf(` and state='%v'`, request.State)
//	}
//	err, total, data := models.NewPool().Pagination(request, whereCondition)
//	if err != nil {
//		log.Logger.Error(err.Error())
//		return statecode.CommonErrServerErr, 0, nil
//	}
//	return statecode.CommonSuccess, total, data
//}
