package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定一个一个密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成token
func ReleaseToken(ID uint) (string, error) {
	//一个token是有过期时间的，这里是一周的过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlarn.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
