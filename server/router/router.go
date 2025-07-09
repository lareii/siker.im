package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lareii/siker.im/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/ping", handlers.Ping)
}
