package server

import (
	userHandler "github.com/vkhoa145/facebook-mini-api/app/modules/users/handler"
	userRepo "github.com/vkhoa145/facebook-mini-api/app/modules/users/repository"
	userUsecase "github.com/vkhoa145/facebook-mini-api/app/modules/users/usecase"
)

func SetupRoutes(server *Server) {
	userRepo := userRepo.NewUserRepo(server.DB)
	userUsecase := userUsecase.NewUserUseCase(userRepo)
	userHandler := userHandler.NewUserHandler(userUsecase)

	api := server.app.Group("/api/v1")
	api.Post("/signup", userHandler.Login())
}
