package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
)

func FavoriteRouter(r fiber.Router, h *handler.FavoriteHandler) {
	r.Get("/", h.GetAll)
	r.Post("/", h.Create)
	r.Delete("/:id", h.Delete)
}
