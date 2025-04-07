package usecase

import "github.com/vkhoa145/facebook-mini-api/app/transaction"

type PostUseCaseInterface interface {
	GetPosts()
}

type PostUseCase struct {
	PostRepo PostUseCaseInterface
	Tx       transaction.TransactionManager
}

func NewPostUseCase(postRepo PostUseCaseInterface, tx transaction.TransactionManager) PostUseCaseInterface {
	return &PostUseCase{
		PostRepo: postRepo,
		Tx:       tx,
	}
}
