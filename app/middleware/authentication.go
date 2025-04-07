package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {

		fmt.Println("authenticationnnn")
		headers := c.GetReqHeaders()

		value, exists := headers["Authorization"]

		if !exists {
			return utils.DataResponseResult(nil, "Authorization Header Must Existed", 401, c)
		}

		_, errClaims := utils.IsValidJwtToken(value[0])
		if errClaims != nil {
			return utils.DataResponseResult(nil, errClaims.Error(), 401, c)
		}

		err := c.Next()
		var isAuthent bool
		if !isAuthent {
			utils.DataResponseResult(nil, "Unauthorized", 401, c)
		}
		return err
	}
}
