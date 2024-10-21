package middleware

import (
	"io"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func NewGormLogger() logger.Interface {
	file, err := os.OpenFile("app/log/development.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Tạo MultiWriter để ghi vào cả stdout và file
	multiWriter := io.MultiWriter(os.Stdout, file)

	gormLogger := logger.New(
		log.New(multiWriter, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Microsecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	return gormLogger
}
