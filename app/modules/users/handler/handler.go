package handler

import (
	user "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
)

type UserHandler struct {
	userUsecase user.UserUseCaseInterface
}

func NewUserHandler(userUsecase user.UserUseCaseInterface) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}
