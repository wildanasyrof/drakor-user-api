// internal/http/middleware/auth.go
package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtpkg "github.com/wildanasyrof/drakor-user-api/pkg/jwt" // alias to avoid name clashes
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
)

func Auth(jwtSvc jwtpkg.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return response.Error(c, fiber.StatusUnauthorized, "Missing bearer token", nil)
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		userID, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid token", err)
		}

		// Stash the uid for handlers/services
		c.Locals("user_id", userID)
		return c.Next()
	}
}
