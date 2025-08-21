package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/di"
	"github.com/wildanasyrof/drakor-user-api/internal/http/middleware"
)

func SetupRouter(app *fiber.App, di *di.DI) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the user api!")
	})

	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })

	AuthRouter(app.Group("/auth"), di.AuthHandler)

	// Protected user endpoints
	me := app.Group("/me")
	me.Use(middleware.Auth(di.JWT)) // <-- inject jwt service into middleware
	UserRouter(me, di.AuthHandler)
}
