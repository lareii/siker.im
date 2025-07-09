package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/lareii/siker.im/database"
	"github.com/lareii/siker.im/router"
	"github.com/lareii/siker.im/utils"
)

func main() {
	utils.LoadEnv()

	database.Setup()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	router.Setup(app)

	log.Fatal(app.Listen(":" + utils.GetEnv("PORT")))
}
