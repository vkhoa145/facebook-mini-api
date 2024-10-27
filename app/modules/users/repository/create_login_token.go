package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

func (r UserRepo) CreateLoginToken(loginToken *models.LoginToken, tx *gorm.DB) (*models.LoginToken, error) {
	createRefreshToken := tx.Table(models.LoginToken{}.TableName()).Create(loginToken)
	if createRefreshToken.Error != nil {
		return nil, createRefreshToken.Error
	}

	return loginToken, nil
}
