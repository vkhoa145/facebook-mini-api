package repository

import "github.com/vkhoa145/facebook-mini-api/app/models"

func (r UserRepo) CreateUser(user *models.User) (*models.User, error) {
	createUser := r.DB.Table(models.User{}.TableName()).Create(&user)
	if createUser.Error != nil {
		return nil, createUser.Error
	}

	return user, nil
}
