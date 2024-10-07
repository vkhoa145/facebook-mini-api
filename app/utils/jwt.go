package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func GenerateJwtToken(user *models.User) string {
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
		fmt.Println("jwt err", err)
	}

	return signedToken
}
