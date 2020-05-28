package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// Jwt json web token
type Jwt interface {
	// Create token
	GenerateToken(claims jwt.Claims) (string, error)

	// 解析token
	ParseWithClaims(token string, claims jwt.Claims) error
}

// JwtAuth jwt
type JwtAuth struct {
	SignKey []byte
}

// New return JwtAuth instance
func New(signKey []byte) Jwt {
	return &JwtAuth{SignKey: signKey}
}

var (
	TokenValidateFailed = errors.New("token validate failed")
)

// CreateToken 生成token
func (jwtAuth JwtAuth) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtAuth.SignKey)
}

// ParseWithClaims 解析token
func (jwtAuth JwtAuth) ParseWithClaims(token string, claims jwt.Claims) error {
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtAuth.SignKey, nil
	})

	if err != nil {
		return  err
	}

	if err := claims.Valid(); err != nil {
		return TokenValidateFailed
	}
	claims = tokenClaims.Claims

	return nil
}
