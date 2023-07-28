package utils

import (
	"cloud-disk/core/define"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateToken 生成Token
func GenerateToken(email string, second int64) (string, error) {
	uc := define.UserClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                          // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	return tokenString, err
}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) { return []byte(define.JwtKey), nil })
	if err != nil {
		return nil, err
	}
	if !claims.Valid { // token过期了
		return uc, errors.New("token已经过期了")
	}
	return uc, err
}
