package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// JWTSecret 是加密签名的密钥
var JWTSecret = []byte("pjfs-df#4216")

// Claims 定义了我们要包含在 JWT 中的数据
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const DefaultExpTime time.Duration = time.Minute * 60

// GenerateToken 生成一个新的 JWT
func GenerateToken(username string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// VerifyToken 验证 JWT 并返回解析后的 Claims
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// 解析并验证 token
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 检查签名算法是否为我们预期的算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 检查 token 是否过期
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
	}

	return claims, nil
}