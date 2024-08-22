package handler

import (
	user "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
)

type UserHandler struct {
	userUsecase user.UserUseCase
}

func NewUserHandler(userUsecase user.UserUseCase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}
