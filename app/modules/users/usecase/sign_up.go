package usecase

import (
	"errors"

	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

func (u UserUseCase) SignUp(user *models.User) (*models.UserResponse, error) {
	email := user.Email
	existingEmail := u.userRepo.CheckExistedEmail(email)
	if existingEmail {
		return nil, errors.New("email is existing")
	}

	tx := u.tx.Begin()
	createdUser, err := u.userRepo.CreateUser(user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	jwt, errJwt := utils.GenerateAccessTokenAndRefreshToken(createdUser)
	if errJwt != nil {
		return nil, errJwt
	}

	createdLoginToken, errLoginToken := u.userRepo.CreateLoginToken(jwt, tx)
	if errLoginToken != nil {
		tx.Rollback()
		return nil, errLoginToken
	}

	tx.Commit()
	userResponse := makeUserResponse(createdUser, createdLoginToken)
	return userResponse, nil
}

func makeUserResponse(user *models.User, jwt *models.JwtResponse) *models.UserResponse {
	return &models.UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Birthday:     user.Birthday,
		Name:         user.Name,
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
	}
}
