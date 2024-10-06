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


		if errorFields := validateSignInParams(payload); errorFields != nil {
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
	hashPassword := utils.HashPassword(payload.Password)
	birthday := utils.ModifyBirthday(int(payload.BirthDay), int(payload.BirthMonth), int(payload.BirthYear))

	fmt.Println("birth day:", birthday)
	user := &models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Birthday: birthday,
		Password: hashPassword,
	}

	return user
}

func validateSignInParams(payload models.SignUpInput) map[string]string {
	errors := make(map[string]string)
	if errorFields := utils.ValidateParams(payload); errorFields != nil {
		errors = handleErrorsMap(errorFields)
	}

	if !utils.IsValidDay(int(payload.BirthDay), int(payload.BirthMonth), int(payload.BirthYear)) {
		errors["BirthDay"] = utils.Locale("en.common_errors.invalid_day")
	}

	return errors
}

func handleErrorsMap(errorsMap map[string]string) map[string]string {
	errors := make(map[string]string)
	for key, errorValue := range errorsMap {
		errors[key] = errorValue
	}

	return errors
}
