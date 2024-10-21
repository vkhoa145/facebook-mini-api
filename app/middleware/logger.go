package middleware

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Logger middleware for logging request details
func Logger() fiber.Handler {
	file, err := os.OpenFile("app/log/development.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	return func(c *fiber.Ctx) error {
		startTime := time.Now()
		startMemory := getMemoryUsage()

		logRequest(file, c, startTime, startMemory)

		// Xử lý request
		err := c.Next()

		endMemory := getMemoryUsage()
		logResponse(file, c, startTime, endMemory, startMemory)

		return err
	}
}

// logRequest ghi lại thông tin bắt đầu request
func logRequest(file *os.File, c *fiber.Ctx, startTime time.Time, startMemory float64) {
	logMsg := fmt.Sprintf("Request started: %s %s at %s, Current Memory: %.2f MB\n",
		c.Method(), c.Path(), startTime.Format(time.RFC3339), startMemory)
	fmt.Fprint(file, logMsg)
	fmt.Print(logMsg)
}

// logResponse ghi lại thông tin kết thúc request
func logResponse(file *os.File, c *fiber.Ctx, startTime time.Time, endMemory, startMemory float64) {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	memoryDiff := endMemory - startMemory

	logMsg := fmt.Sprintf("Request ended: %s %s at %s, Duration: %v, Memory usage: %.2f MB, Memory diff: %.2f MB\n",
		c.Method(), c.Path(), endTime.Format(time.RFC3339), duration, endMemory, memoryDiff)
	fmt.Fprint(file, logMsg)
	fmt.Print(logMsg)
}

// getMemoryUsage trả về lượng memory đang sử dụng
func getMemoryUsage() float64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return float64(memStats.Alloc) / 1024 / 1024
}
