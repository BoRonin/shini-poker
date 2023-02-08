package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	secret = "shini"
)

func GenerateJwt(id int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(id),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	return claims.SignedString([]byte(secret))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil

}
