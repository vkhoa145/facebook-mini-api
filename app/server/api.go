package server

import (
	userHandler "github.com/vkhoa145/facebook-mini-api/app/modules/users/handler"
	userRepo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	userUsecase "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
	"github.com/vkhoa145/facebook-mini-api/app/queries"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
)

func SetupRoutes(server *Server) {
	queries := queries.NewQueries(server.DB)
	transactionManager := transaction.NewTransactionManager(server.DB)
	userRepo := userRepo.NewUserRepo(server.DB, queries)
	userUsecase := userUsecase.NewUserUseCase(userRepo, *transactionManager)
	userHandler := userHandler.NewUserHandler(userUsecase)

	api := server.app.Group("/api/v1")
	api.Post("/signup", userHandler.Login())
}
