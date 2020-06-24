package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"g1/auth/model"
)

var jwtKey = []byte("a_secret_value")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	// sha256  -> 私钥签名 + 明文
	// 公钥解密 -> sha256  + 明文sha256 与sha256比对
	// token hash
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}