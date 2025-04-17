package protocol

type LoginRequest struct {
	LoginName string `form:"loginName" binding:"required"`
	Password  string `form:"password" binding:"required"`
}
type LoginResponse struct {
}
