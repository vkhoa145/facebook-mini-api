package repository

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"gorm.io/gorm"
)

func (r UserRepo) CreateUser(user *models.User, tx *gorm.DB) (*models.User, error) {
	createUser := tx.Table(models.User{}.TableName()).Create(&user)
	if createUser.Error != nil {
		return nil, createUser.Error
	}

	return user, nil
}
