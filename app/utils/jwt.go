package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func GenerateAccessTokenAndRefreshToken(user *models.User) (*models.JwtResponse, error) {
	accessToken, errorAccessToken := GenerateAccessToken(user)
	if errorAccessToken != nil {
		return nil, errorAccessToken
	}

	refreshToken, errorRefreshToken := GenerateRefreshToken()
	if errorRefreshToken != nil {
		return nil, errorRefreshToken
	}

	jwtResponse := &models.JwtResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return jwtResponse, nil
}

func GenerateAccessToken(user *models.User) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := models.Claims{
		UserId: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken() (string, error) {
	refreshToken, err := GenerateRandomString(32)
	if err != nil {
		return "", nil
	}
	return refreshToken, nil
}
