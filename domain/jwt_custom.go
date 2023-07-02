package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims struct for access token
type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

// JwtCustomRefreshClaims struct for refresh token
type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
