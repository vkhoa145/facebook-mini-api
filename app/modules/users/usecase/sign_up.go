package usecase

import (
	"errors"
	"time"

	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

type Result struct {
	CreateUserResponse *models.UserResponse
	VerifyCode         *models.VerificationCode
	ErrorResponse      error
}

func (u UserUseCase) SignUp(user *models.User) (*Result, error) {
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
		UserID:       createdUser.ID,
		RefreshToken: jwt.RefreshToken,
	}

	_, errLoginToken := u.UserRepo.CreateLoginToken(loginToken, tx)
	if errLoginToken != nil {
		tx.Rollback()
		return nil, errLoginToken
	}

	verifycationCode := &models.VerificationCode{
		UserID:           createdUser.ID,
		VerifycationType: "email",
		Code:             utils.GenerateVerifyCode(),
		ExpiredAt:        time.Now().Add(15 * time.Minute),
	}

	_, errVerifyCode := u.UserRepo.CreateVerificationCode(verifycationCode, tx)
	if errVerifyCode != nil {
		tx.Rollback()
		return nil, errVerifyCode
	}

	tx.Commit()

	userResponse := makeUserResponse(createdUser, jwt.AccessToken, jwt.RefreshToken)
	result := &Result{
		CreateUserResponse: userResponse,
		VerifyCode:         verifycationCode,
	}
	return result, nil
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
