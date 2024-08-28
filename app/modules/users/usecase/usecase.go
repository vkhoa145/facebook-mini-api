package usecase

import (
	repo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
)

type UserUseCaseInterface interface {
}

type UserUseCase struct {
	userRepo repo.UserRepoInterface
}

func NewUserUseCase(userRepo repo.UserRepoInterface) UserUseCaseInterface {
	return &UserUseCase{userRepo: userRepo}
}
