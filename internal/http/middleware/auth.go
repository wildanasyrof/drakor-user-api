// internal/http/middleware/auth.go
package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	jwtpkg "github.com/wildanasyrof/drakor-user-api/pkg/jwt"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
)

func Auth(jwtSvc jwtpkg.JWTService, log logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return response.Error(c, fiber.StatusUnauthorized, "Missing bearer token", nil)
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		userIDStr, err := jwtSvc.ValidateToken(tokenStr) // returns user id as string (e.g., JWT sub)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid token", err)
		}

		// Parse to UUID once, here.
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "User ID in token is not a valid UUID", nil)
		}

		// Stash the typed UUID for handlers/services
		c.Locals("user_id", userID)
		return c.Next()
	}
}
