package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vkhoa145/facebook-mini-api/config"
	db "github.com/vkhoa145/facebook-mini-api/db"
	"gorm.io/gorm"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	DB     *gorm.DB
}

func NewServer(fiber *fiber.App, cfg *config.Config) *Server {
	return &Server{
		app:    fiber,
		config: cfg,
		DB:     db.NewDb(cfg),
	}
}

func (server *Server) Start() error {
	server.app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
	}))

	SetupRoutes(server)
	return server.app.Listen(":" + server.config.HTTP.Port)
}
