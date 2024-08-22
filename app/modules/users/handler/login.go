package handler

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"status": "oke"})
	}
}
