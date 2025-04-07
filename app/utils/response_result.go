package utils

import "github.com/gofiber/fiber/v2"

func DataResponseResult(result interface{}, failure interface{}, status int, ctx *fiber.Ctx) error {
	ctx.Status(status)
	return ctx.JSON(&fiber.Map{"status": status, "result": result, "error": failure})
}


