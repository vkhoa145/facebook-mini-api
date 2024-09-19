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
			fmt.Println("error fields:", errorFields)
			return utils.DataResponseResult(nil, errorFields, 400, ctx)
		}

		createdUser, err := h.userUsecase.SignUp(payload)

		if err != nil {
			return utils.DataResponseResult(nil, err.Error(), 400, ctx)
		}

		return utils.DataResponseResult(createdUser, nil, 200, ctx)
	}
}
