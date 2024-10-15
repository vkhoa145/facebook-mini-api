package server

import (
	userHandler "github.com/vkhoa145/facebook-mini-api/app/modules/users/handler"
	userRepo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	userUsecase "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
)

func SetupRoutes(server *Server) {
	transactionManager := transaction.NewTransactionManager(server.DB)
	userRepo := userRepo.NewUserRepo(server.DB, server.DB.Begin())
	userUsecase := userUsecase.NewUserUseCase(userRepo, *transactionManager)
	userHandler := userHandler.NewUserHandler(userUsecase)

	api := server.app.Group("/api/v1")
	api.Post("/signup", userHandler.Login())
}
