package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"smart-route/pkg/config"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Address string `json:"address"`
	jwt.RegisteredClaims
}

type AuthService struct {
	config *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		config: cfg,
	}
}

// GenerateNonce 生成随机 nonce
func (a *AuthService) GenerateNonce() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// VerifySignature 验证钱包签名
func (a *AuthService) VerifySignature(address, message, signature string) bool {
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
func (a *AuthService) GenerateToken(address string) (string, error) {
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
func (a *AuthService) AuthMiddleware() gin.HandlerFunc {
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
