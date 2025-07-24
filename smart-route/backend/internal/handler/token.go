package handler

import (
	"net/http"
	"smart-route/internal/model"
	"strconv"

	"smart-route/internal/api"
	"smart-route/internal/service"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenService *service.TokenService
}

func NewTokenHandler(tokenService *service.TokenService) *TokenHandler {
	return &TokenHandler{
		tokenService: tokenService,
	}
}

// ListTokens @Summary 查询 Token 列表
// @Tags token
// @Accept json
// @Produce json
// @Param chain query string false "链名称"
// @Param chainId query int false "链ID"
// @Param name query string false "Token名称"
// @Param symbol query string false "Token符号"
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} dto.TokenListResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /token/list [get]
func (h *TokenHandler) ListTokens(c *gin.Context) {
	chain := c.Query("chain")
	chainIdStr := c.Query("chainId")
	name := c.Query("name")
	symbol := c.Query("symbol")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	//var req api.TokenListRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
	//	return
	//}

	var chainId *int
	if chainIdStr != "" {
		if v, err := strconv.Atoi(chainIdStr); err == nil {
			chainId = &v
		} else {
			c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: "invalid chainId"})
			return
		}
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	tokens, total, err := h.tokenService.ListTokens(
		c.Request.Context(),
		chain,
		chainId,
		name,
		symbol,
		page,
		pageSize,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Failed to query token list"})
		return
	}

	resp := api.TokenListResponse{
		Tokens:   tokens,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	c.JSON(http.StatusOK, resp)
}

// CreateToken @Summary 录入 Token
// @Tags admin
// @Accept json
// @Produce json
// @Param request body dto.AssetCreateRequest true "Token参数"
// @Success 200 {object} model.Token
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /admin/token/create [post]
func (h *TokenHandler) CreateToken(c *gin.Context) {
	var req api.TokenCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	token := model.Token{
		Symbol:   req.Symbol,
		Name:     req.Name,
		Address:  req.Address,
		Decimals: req.Decimals,
		Chain:    req.Chain,
		ChainId:  req.ChainId,
		IsStable: req.IsStable,
	}
	if err := h.tokenService.CreateToken(c.Request.Context(), &token); err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Error: "Failed to create token"})
		return
	}
	c.JSON(http.StatusOK, token)
}
