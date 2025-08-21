package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
)

func UserRouter(r fiber.Router, h *handler.AuthHandler) {
	r.Post("/refresh", h.Refresh)
	r.Post("/logout", h.Logout)
}
