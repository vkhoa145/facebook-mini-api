package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	setting := fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Facebook Mini",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	app := fiber.New(setting)
	return app
}
