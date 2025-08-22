// internal/http/middleware/request_logger.go
package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
)

func LoggerMiddleware(log logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Basic request info
		method := c.Method()
		path := c.OriginalURL()
		// req := c.Request().Body()
		ip := c.IP()
		ua := string(c.Request().Header.UserAgent())

		// Continue down the chain
		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()

		uid, _ := c.Locals("user_id").(string) // set by your Auth middleware for /me routes

		// If handlers returned an error, log it at error level as well
		if err != nil {
			log.Error(err, "request error")
		}

		// Structured single-line summary
		log.Info(fmt.Sprintf(
			`%s %s -> %d in %s ip=%s uid=%s ua=%q`,
			method, path, status, latency, ip, uid, ua,
		))

		return err
	}
}
