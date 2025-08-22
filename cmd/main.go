package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/internal/di"
	"github.com/wildanasyrof/drakor-user-api/internal/http/router"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	di := di.InitDI(cfg)
	router.SetupRouter(app, di, cfg)

	app.Listen(":" + cfg.Server.Port)
}
