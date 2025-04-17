package protocol

type OrderDetailRequest struct {
	ID int `json:"id" binding:"required"`
}
