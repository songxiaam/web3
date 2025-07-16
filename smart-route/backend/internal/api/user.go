package api

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	Address   string `json:"address" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	Token   string `json:"token"`
	Address string `json:"address"`
}

type UserInfoRequest struct {
	Address string `json:"address"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
}
