package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"smart-route/internal/api"
	"smart-route/internal/auth"
	"smart-route/internal/service"
)

type UserHandler struct {
	jtwAuth     *auth.JwtAuth
	userService *service.UserService
}

func NewUserHandler(jwtAuth *auth.JwtAuth, userService *service.UserService) *UserHandler {
	return &UserHandler{
		jtwAuth:     jwtAuth,
		userService: userService,
	}
}

// Login @Summary 用户登录
// @Description 使用钱包签名进行登录
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "登录请求参数"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req api.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	// 校验签名
	if !h.jtwAuth.VerifySignature(req.Address, req.Message, req.Signature) {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid signature"})
		return
	}

	//userObj, err := dataUser.GetUserByAddress(c.Request.Context(), uh.db, req.Address)
	userModel, err := h.userService.GetUserByAddress(c, req.Address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//userModel, err = dataUser.CreateUser(c.Request.Context(), uh.db, req.Address)

			userModel, err = h.userService.CreateUser(c, req.Address)

			if err != nil {
				c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Failed to create user"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Database error"})
			return
		}
	}

	// 生成 JWT
	token, err := h.jtwAuth.GenerateToken(userModel.ID.String(), userModel.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, api.UserLoginResponse{
		Token:   token,
		Address: userModel.Address,
	})
}

// GetProfile @Summary 获取用户信息
// @Description 获取当前登录用户的信息
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ProfileResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /api/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	var req api.UserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := h.userService.GetUserByAddress(c, req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Database error"})
		return
	}
	c.JSON(http.StatusOK, api.UserInfoResponse{
		ID:       user.ID.String(),
		Nickname: user.Nickname,
		Address:  user.Address,
		Avatar:   user.Avatar,
		Status:   user.Status,
	})
}
