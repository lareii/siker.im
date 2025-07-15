package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/lareii/siker.im/internal/config"
)

func setupRoutes(app *fiber.App, deps *dependencies, cfg *config.Config) {
	setupGlobalMiddleware(app, cfg)
	setupAPIRoutes(app, deps)
}

func setupGlobalMiddleware(app *fiber.App, cfg *config.Config) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.App.AllowedOrigins,
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Cf-Turnstile-Token"},
	}))
}

func setupAPIRoutes(app *fiber.App, deps *dependencies) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK!",
		})
	})

	app.Post("/urls", deps.URLHandler.CreateURL, deps.Middleware.RateLimiter.Middleware(), deps.Middleware.Turnstile.Verify())
	app.Get("/urls/:param", deps.URLHandler.GetURL)
	app.Get("/redirect/:slug", deps.URLHandler.RedirectURL)
}
