package handler

import (
	"net/http"
	"smart-route/internal/api"
	"smart-route/internal/auth"
	"smart-route/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *service.AdminService
	jtwAuth      *auth.JwtAuth
}

func NewAdminHandler(jwtAuth *auth.JwtAuth, adminService *service.AdminService) *AdminHandler {
	return &AdminHandler{
		jtwAuth:      jwtAuth,
		adminService: adminService}
}

// Login @Summary 管理员登录
// @Tags admin
// @Accept json
// @Produce json
// @Param request body dto.AdminLoginRequest true "登录参数"
// @Success 200 {object} dto.AdminLoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /admin/login [post]
func (h *AdminHandler) Login(c *gin.Context) {
	var req api.AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	admin, err := h.adminService.CheckAdmin(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: err.Error()})
		return
	}

	// 这里可生成 JWT
	token, err := h.jtwAuth.GenerateTokenAdmin(admin.ID.String())

	c.JSON(http.StatusOK, api.AdminLoginResponse{
		Token:    token,
		Username: admin.Username,
		Role:     admin.Role,
	})
}
