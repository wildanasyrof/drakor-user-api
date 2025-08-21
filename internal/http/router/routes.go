package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/di"
)

func SetupRouter(app *fiber.App, di *di.DI) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the user api!")
	})

}
