package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/facebook-mini-api/config"
)

type Server struct {
	Fiber *fiber.App
	cfg   *config.Config
}

func NewServer(fiber *fiber.App, cfg *config.Config) *Server {
	return &Server{
		Fiber: fiber,
		cfg:   cfg,
	}
}



