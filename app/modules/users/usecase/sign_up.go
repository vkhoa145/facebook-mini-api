package usecase

import (
	"errors"
	"fmt"

	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

func (u UserUseCase) SignUp(user *models.User) (*models.UserResponse, error) {
	email := user.Email
	existingEmail := u.userRepo.CheckExistedEmail(email)
	if existingEmail {
		return nil, errors.New("email is existing")
	}

	createdUser, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	jwt := utils.GenerateJwtToken(createdUser)
	userResponse := makeUserResponse(createdUser, jwt)
	fmt.Println("user response", userResponse)
	return userResponse, nil
}

func makeUserResponse(user *models.User, jwt string) *models.UserResponse {
	return &models.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		Birthday:    user.Birthday,
		Name:        user.Name,
		AccessToken: jwt,
	}
}
