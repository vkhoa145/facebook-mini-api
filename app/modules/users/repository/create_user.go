package repository

import "github.com/vkhoa145/facebook-mini-api/app/models"

func (r UserRepo) CreateUser(payload *models.SignUpInput) (*models.User, error) {
	user := &models.User{
		Email: payload.Email,
		Name:  payload.Name,
	}

	createUser := r.DB.Table(models.User{}.TableName()).Create(user)

	if createUser.Error != nil {
		return nil, createUser.Error
	}

	return user, nil
}
