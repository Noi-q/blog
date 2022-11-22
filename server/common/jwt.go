package common

import (
	"blog-admin/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成token
func ReleaseToken(admin models.Admin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: admin.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // token过期时间
			IssuedAt:  time.Now().Unix(),     // token发放时间
			Issuer:    "Noi-q",               // token发放者
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

// 验证token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
