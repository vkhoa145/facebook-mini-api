package usecase

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	repo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
)

type UserUseCaseInterface interface {
	SignUp(payload *models.User) (*models.UserResponse, error)
}

type UserUseCase struct {
	userRepo repo.UserRepoInterface
}

func NewUserUseCase(userRepo repo.UserRepoInterface) UserUseCaseInterface {
	return &UserUseCase{userRepo: userRepo}
}
