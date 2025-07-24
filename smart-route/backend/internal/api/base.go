package api

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid signature"`
}
