package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	utils "github.com/vkhoa145/facebook-mini-api/app/utils"
)

func (h *UserHandler) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignUpInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			return utils.DataResponseResult(nil, err.Error(), 400, ctx)
		}

		if errorFields := utils.ValidateParams(payload); errorFields != nil {
			println("error field", errorFields)
			return utils.DataResponseResult(nil, errorFields, 400, ctx)
		}

		user := modifyUserParams(&payload)
		createdUser, err := h.userUsecase.SignUp(user)

		if err != nil {
			return utils.DataResponseResult(nil, err.Error(), 400, ctx)
		}

		return utils.DataResponseResult(createdUser, nil, 200, ctx)
	}
}

func modifyUserParams(payload *models.SignUpInput) *models.User {
	BirthDay := fmt.Sprintf("%d/%d/%d", payload.BirthDay, payload.BirthMonth, payload.BirthYear)
	hashPassword := utils.HashPassword(payload.Password)
	fmt.Println("password:", hashPassword)
	fmt.Println("BirthDay:", BirthDay)
	fmt.Println("BirthDay:", payload.BirthDay)
	fmt.Println("BirthMonth:", payload.BirthMonth)
	fmt.Println("BirthYear:", payload.BirthYear)
	user := &models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Birthday: BirthDay,
		Password: hashPassword,
	}

	return user
}
