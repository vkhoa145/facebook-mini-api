package usecase

import (
	"errors"

	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

func (u UserUseCase) SignUp(user *models.User) (*models.UserResponse, error) {
	email := user.Email
	existingEmail := u.UserRepo.CheckExistedEmail(email)
	if existingEmail {
		return nil, errors.New("email is existing")
	}

	tx := u.Tx.Begin()
	createdUser, err := u.UserRepo.CreateUser(user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	jwt, errJwt := utils.GenerateAccessTokenAndRefreshToken(createdUser)
	if errJwt != nil {
		return nil, errJwt
	}

	loginToken := &models.LoginToken{
		UserID: createdUser.ID,
	}

	_, errLoginToken := u.UserRepo.CreateLoginToken(loginToken, tx)
	if errLoginToken != nil {
		tx.Rollback()
		return nil, errLoginToken
	}

	tx.Commit()
	userResponse := makeUserResponse(createdUser, jwt.AccessToken, jwt.RefreshToken)
	return userResponse, nil
}

func makeUserResponse(user *models.User, accessToken string, refreshToken string) *models.UserResponse {
	return &models.UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Birthday:     user.Birthday,
		Name:         user.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
