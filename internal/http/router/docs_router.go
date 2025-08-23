package router

import (
	"github.com/gofiber/fiber/v2"
)

func DocsRouter(r fiber.Router) {
	// Serve the raw YAML file from disk
	r.Static("/openapi.yaml", "./openapi.yaml")

	// ReDoc
	r.Get("/docs", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`<!doctype html>
<html>
  <head>
    <meta charset="utf-8"/>
    <title>API Docs (ReDoc)</title>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
  </head>
  <body>
    <redoc spec-url="/openapi.yaml"></redoc>
  </body>
</html>`)
	})

	// Swagger UI
	r.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`<!doctype html>
<html>
  <head>
    <meta charset="utf-8"/>
    <title>API Docs (Swagger UI)</title>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui.css">
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui-bundle.js"></script>
    <script>
      window.ui = SwaggerUIBundle({
        url: "/openapi.yaml",
        dom_id: "#swagger-ui",
        deepLinking: true
      });
    </script>
  </body>
</html>`)
	})
}
