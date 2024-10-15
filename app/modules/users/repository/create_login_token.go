package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

func (r UserRepo) CreateLoginToken(jwt *models.JwtResponse, tx *gorm.DB) (*models.JwtResponse, error) {
	LoginToken := &models.LoginToken{
		RefreshToken: jwt.RefreshToken,
	}

	createRefreshToken := tx.Table(models.LoginToken{}.TableName()).Create(LoginToken)
	if createRefreshToken.Error != nil {
		return nil, createRefreshToken.Error
	}

	return jwt, nil
}
