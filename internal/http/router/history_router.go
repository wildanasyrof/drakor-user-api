package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
)

func HistoryRouter(r fiber.Router, handler *handler.HistoryHandler) {
	r.Post("/", handler.Create)
	r.Delete("/:id", handler.Delete)
	r.Get("/", handler.GetAll)
	r.Put("/:id", handler.Update)
}
