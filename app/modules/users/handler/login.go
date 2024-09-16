package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/app/models"
)

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.SignUpInput{}
		if err := c.BodyParser(payload); err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{"status": 400, "message": err.Error()})
		}

		createdUser, err := h.userUsecase.SignUp(payload)

		if err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{"status": 400, "error": err.Error()})
		}

		c.Status(200)
		return c.JSON(&fiber.Map{"status": 200, "data": createdUser})
	}
}
