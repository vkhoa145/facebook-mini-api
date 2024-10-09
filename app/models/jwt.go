package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type JwtResponse struct {
	AccessToken string
	RefreshToken string
}
