package protocol

type ProductListRequest struct {
	StartIndex int `json:"startIndex" binding:"required"`
	PageSize   int `json:"pageSize" binding:"required"`
}
