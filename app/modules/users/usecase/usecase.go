package usecase

import (
	"github.com/vkhoa145/facebook-mini-api/app/models"
	repo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
)

type UserUseCaseInterface interface {
	SignUp(payload *models.User) (*Result, error)
}

type UserUseCase struct {
	UserRepo repo.UserRepoInterface
	Tx       transaction.TransactionManager
}

func NewUserUseCase(userRepo repo.UserRepoInterface, tx transaction.TransactionManager) UserUseCaseInterface {
	return &UserUseCase{
		UserRepo: userRepo,
		Tx:       tx,
	}
}
