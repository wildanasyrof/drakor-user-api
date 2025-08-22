package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
)

func UserRouter(r fiber.Router, a *handler.AuthHandler, u *handler.UserHandler) {
	r.Post("/refresh", a.Refresh)
	r.Post("/logout", a.Logout)
	r.Get("/", u.GetProfile)
	r.Put("/", u.UpdateProfile)
	r.Put("/avatar", u.UpdateAvatar)
}
