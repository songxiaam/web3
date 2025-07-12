package api

import (
	"net/http"

	"smart-route/internal/auth"
	"smart-route/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	authService *auth.AuthService
	config      *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		authService: auth.NewAuthService(cfg),
		config:      cfg,
	}
}

// NewRouter 创建 gin 路由
func NewRouter(cfg *config.Config) *gin.Engine {
	h := NewHandler(cfg)
	r := gin.Default()

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 钱包登录相关路由
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/nonce", h.getNonce)
		authGroup.POST("/login", h.walletLogin)
	}

	// 需要鉴权的路由
	protected := r.Group("/api")
	protected.Use(h.authService.AuthMiddleware())
	{
		protected.GET("/profile", h.getProfile)
	}

	return r
}

// @Summary 获取随机 nonce
// @Description 获取用于钱包签名的随机 nonce
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "返回随机 nonce"
// @Router /auth/nonce [get]
func (h *Handler) getNonce(c *gin.Context) {
	nonce := h.authService.GenerateNonce()
	c.JSON(http.StatusOK, gin.H{
		"nonce": nonce,
	})
}

// @Summary 钱包登录
// @Description 使用钱包签名进行登录
// @Tags auth
// @Accept json
// @Produce json
// @Param request body WalletLoginRequest true "登录请求参数"
// @Success 200 {object} WalletLoginResponse "登录成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "签名验证失败"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /auth/login [post]
func (h *Handler) walletLogin(c *gin.Context) {
	var req WalletLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证签名
	if !h.authService.VerifySignature(req.Address, req.Message, req.Signature) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	// 生成 JWT token
	token, err := h.authService.GenerateToken(req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, WalletLoginResponse{
		Token:   token,
		Address: req.Address,
	})
}

// @Summary 获取用户信息
// @Description 获取当前登录用户的信息
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ProfileResponse "用户信息"
// @Failure 401 {object} ErrorResponse "未授权"
// @Router /api/profile [get]
func (h *Handler) getProfile(c *gin.Context) {
	address, _ := c.Get("address")
	c.JSON(http.StatusOK, ProfileResponse{
		Address: address.(string),
		Message: "Profile retrieved successfully",
	})
}

// WalletLoginRequest 钱包登录请求
type WalletLoginRequest struct {
	Address   string `json:"address" binding:"required" example:"0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6"`
	Message   string `json:"message" binding:"required" example:"Login to Smart Route"`
	Signature string `json:"signature" binding:"required" example:"0x1234567890abcdef..."`
}

// WalletLoginResponse 钱包登录响应
type WalletLoginResponse struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6"`
}

// ProfileResponse 用户信息响应
type ProfileResponse struct {
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6"`
	Message string `json:"message" example:"Profile retrieved successfully"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid signature"`
}
