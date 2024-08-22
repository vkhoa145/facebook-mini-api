package usecase

import "github.com/gofiber/fiber/v2"

type userUseCaseInterface interface {
	SignIn(ctx *fiber.Ctx) error
}

type UserUseCase struct {
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}
