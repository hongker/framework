package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func TestJwtAuth_CreateToken(t *testing.T) {
	auth := New([]byte("123456"))
	token, err := auth.GenerateToken(&CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600,
			Issuer:    "test",
		},
		UserId:         1,
	})

	assert.Nil(t, err)
	fmt.Println(token)

	claims := new(CustomClaims)
	err = auth.ParseWithClaims(token, claims)
	assert.Nil(t, err)
	fmt.Println(claims.UserId)
}
