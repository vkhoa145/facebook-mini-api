package handler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/app/models"
	"github.com/vkhoa145/facebook-mini-api/app/services"
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
		result, err := h.userUsecase.SignUp(user)

		if err != nil {
			return utils.DataResponseResult(nil, err.Error(), 400, ctx)
		}

		if user.Email != "" {
			content := makeVerificationEmailContent(user, result.VerifyCode)
			subject := "Verification Account Email"
			services.SendEmail("khoavodang1451997@gmail.com", content, subject)
		} else {
			fmt.Println("helloooo")
		}

		return utils.DataResponseResult(result.CreateUserResponse, nil, 200, ctx)
	}
}

func modifyUserParams(payload *models.SignUpInput) *models.User {
	hashPassword := utils.HashPassword(payload.Password)
	birthday := utils.ModifyBirthday(int(payload.BirthDay), int(payload.BirthMonth), int(payload.BirthYear))
	emailVerifyCode := makeVerifyCode(payload.Email)
	phoneVerifyCode := makeVerifyCode(payload.Phone)
	user := &models.User{
		Email:           payload.Email,
		Name:            payload.Name,
		Birthday:        birthday,
		Password:        hashPassword,
		EmailVerifycode: emailVerifyCode,
		PhoneVerifycode: phoneVerifyCode,
		VerifycatedAt:   time.Now(),
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

	var errorKeys []string
	for key := range errors {
		errorKeys = append(errorKeys, key)
	}

	if len(errorKeys) == 0 {
		return nil
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

func makeVerifyCode(params string) string {
	if params == "" {
		return ""
	}

	verifyCode, err := generateRandomString(100)
	if err != nil {
		return ""
	}

	return verifyCode
}

func makeVerificationEmailContent(user *models.User, verifyCode *models.VerificationCode) string {
	verificationAt := utils.FormatDateTime(verifyCode.ExpiredAt)
	return fmt.Sprintf(`<h1>Thanks for your register, here is verification code</h1>
	<p>This email was sent to <strong>%s</strong></p> 
	<p>Here is your code <strong>%s</strong></p> 
	<p>This code will be expired at: <strong>%s</strong> "</p>`, user.Email, verifyCode.Code, verificationAt)
}

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:',.<>?/"
	bytes := make([]byte, length)
	for i := range bytes {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		bytes[i] = charset[randomByte.Int64()]
	}
	return string(bytes), nil
}
