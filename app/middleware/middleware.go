package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
)

type MiddlewareInterface interface {
	NewGormLogger() logger.Interface
	Logger() fiber.Handler
}

type Middlewrare struct {
	File *os.File
}

func NewMiddleware(file *os.File) *Middlewrare {
	return &Middlewrare{
		File: file,
	}
}
