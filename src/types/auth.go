package types

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type Public struct {
	Path string
	Method string
}