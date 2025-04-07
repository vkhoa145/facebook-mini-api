package server

import (
	postHandler "github.com/vkhoa145/facebook-mini-api/app/modules/posts/handler"
	postRepo "github.com/vkhoa145/facebook-mini-api/app/modules/posts/repository"
	postUsecase "github.com/vkhoa145/facebook-mini-api/app/modules/posts/usecase"
	userHandler "github.com/vkhoa145/facebook-mini-api/app/modules/users/handler"
	userRepo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	userUsecase "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
	"github.com/vkhoa145/facebook-mini-api/app/queries"
	"github.com/vkhoa145/facebook-mini-api/app/transaction"
)

func SetupRoutes(server *Server) {
	queries := queries.NewQueries(server.DB)
	transactionManager := transaction.NewTransactionManager(server.DB)

	// /signup
	userRepo := userRepo.NewUserRepo(server.DB, queries)
	userUsecase := userUsecase.NewUserUseCase(userRepo, *transactionManager)
	userHandler := userHandler.NewUserHandler(userUsecase)

	// /signin

	// /posts
	postRepo := postRepo.NewPostRepo(server.DB, queries)
	postUsecase := postUsecase.NewPostUseCase(postRepo, *transactionManager)
	postHandler := postHandler.NewPostHandler(postUsecase)

	api := server.app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/signup", userHandler.Login())
	api.Get("/posts", postHandler.GetPosts())
}
