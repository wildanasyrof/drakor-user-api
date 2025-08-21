package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
)

func AuthRouter(r fiber.Router, h *handler.AuthHandler) {
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
}
