package utils

import (
	"app/src/types"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(id, secret string) (string, error){
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &types.Claims {
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJwtToken(token, secret string) (*types.Claims, error){
	claims := &types.Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error){
		log.Println(token.Claims)
		return []byte(secret), nil
	})
	return claims, err
}