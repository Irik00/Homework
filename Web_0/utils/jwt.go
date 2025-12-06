package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret key")

func MakeToken(username string, expireTime time.Time) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"user":     expireTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
