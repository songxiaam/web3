package auth

import (
	"crypto/rand"
	"encoding/hex"
	"smart-route/pkg/data/entity"
	"time"

	"gorm.io/gorm"

	"smart-route/pkg/config"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuth struct {
	config *config.Config
	db     *gorm.DB
}

func NewJwtAuth(cfg *config.Config, db *gorm.DB) *JwtAuth {
	return &JwtAuth{
		config: cfg,
		db:     db,
	}
}

type Claims struct {
	Address string `json:"address"`
	jwt.RegisteredClaims
}

type AdminClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateNonce 生成随机 nonce
func (a *JwtAuth) GenerateNonce() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// VerifySignature 验证钱包签名
func (a *JwtAuth) VerifySignature(address, message, signature string) bool {
	// 解码签名
	sig, err := hex.DecodeString(signature[2:]) // 移除 0x 前缀
	if err != nil {
		return false
	}

	// 验证签名长度
	if len(sig) != 65 {
		return false
	}

	// 恢复公钥
	hash := crypto.Keccak256Hash([]byte(message))
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}

	// 从公钥恢复地址
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr.Hex() == address
}

// GenerateToken 生成 JWT token
func (a *JwtAuth) GenerateToken(userId, address string) (string, error) {
	claims := Claims{
		Address: address,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.config.JWT.ExpiresHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.config.JWT.Secret))
}

// AuthMiddleware JWT 鉴权中间件
func (a *JwtAuth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// 移除 Bearer 前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.config.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("address", claims.Address)
		}

		c.Next()
	}
}

// GenerateTokenAdmin 生成 JWT token
func (a *JwtAuth) GenerateTokenAdmin(userId string) (string, error) {
	claims := AdminClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.config.JWT.ExpiresHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.config.JWT.Secret))
}

// AdminAuthMiddleware 管理员用户名密码鉴权中间件
func (a *JwtAuth) AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(400, gin.H{"error": "Username and password required"})
			c.Abort()
			return
		}
		var adminUser entity.Admin
		if err := a.db.Where("username = ?", login.Username).First(&adminUser).Error; err != nil {
			c.JSON(401, gin.H{"error": "Admin user not found"})
			c.Abort()
			return
		}
		if adminUser.Password != login.Password {
			c.JSON(401, gin.H{"error": "Invalid admin credentials"})
			c.Abort()
			return
		}

		// 假设管理员用户名和密码存储在配置中
		if login.Username != adminUser.Username || login.Password != adminUser.Password {
			c.JSON(401, gin.H{"error": "Invalid admin credentials"})
			c.Abort()
			return
		}

		// 认证通过，设置管理员标识
		c.Set("is_admin", true)
		c.Next()
	}
}
