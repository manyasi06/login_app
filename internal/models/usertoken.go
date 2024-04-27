package models

import "github.com/golang-jwt/jwt/v5"

type JwtTokenCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
