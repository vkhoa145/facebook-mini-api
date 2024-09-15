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

		if c.Params("name") == "" {
			return c.JSON(&fiber.Map{"status": "400", "message": "field name is blank"})
		}
		return c.JSON(&fiber.Map{"status": "oke"})
	}
}
