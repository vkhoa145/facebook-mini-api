package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/app/utils"
)

func (h *PostHandler) GetPosts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return utils.DataResponseResult("oke", nil, 200, ctx)
	}
}
