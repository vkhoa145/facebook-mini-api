package usecase

import (
	"errors"

	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func (u UserUseCase) SignUp(payload models.SignUpInput) (*models.UserResponse, error) {
	email := payload.Email
	existingEmail := u.userRepo.CheckExistedEmail(email)
	if existingEmail {
		return nil, errors.New("email is existing")
	}

	createdUser, err := u.userRepo.CreateUser(&payload)
	if err != nil {
		return nil, err
	}

	userResponse := makeUserResponse(createdUser)
	return userResponse, nil
}

func makeUserResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Birthday: user.Birthday,
	}
}
