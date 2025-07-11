package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lareii/siker.im/internal/config"
	"github.com/lareii/siker.im/internal/database"
	"github.com/lareii/siker.im/internal/handlers"
	"github.com/lareii/siker.im/internal/middleware"
	"github.com/lareii/siker.im/internal/repository"
	"github.com/lareii/siker.im/internal/services"
	"github.com/lareii/siker.im/pkg/logger"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	logger := logger.New(cfg.LogLevel)
	defer logger.Sync()

	db, err := database.NewMongoDB(cfg.Database.URI, cfg.Database.Name)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Disconnect(context.Background())
	logger.Info("Connected to MongoDB", zap.String("uri", cfg.Database.URI))

	redis, err := database.NewRedis(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		logger.Fatal("Failed to connect to Redis", zap.Error(err))
	}
	defer redis.Close()
	logger.Info("Connected to Redis", zap.String("host", cfg.Redis.Host), zap.String("port", cfg.Redis.Port))

	urlRepo := repository.NewURLRepository(db)
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService, logger)
	rateLimiter := middleware.NewRateLimiter(redis, &cfg.RateLimit, logger)

	app := fiber.New(fiber.Config{
		AppName:      "siker.im",
		ServerHeader: "Fiber",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			logger.Error("Request error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.App.AllowedOrigins},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	app.Use(rateLimiter.Middleware())

	setupRoutes(app, urlHandler)

	go func() {
		logger.Info("Server starting on port " + cfg.Server.Port)
		if err := app.Listen(":" + cfg.Server.Port); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited")
}

func setupRoutes(app *fiber.App, urlHandler *handlers.URLHandler) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK!",
		})
	})

	app.Post("/urls", urlHandler.CreateURL)
	app.Get("/urls/:param", urlHandler.GetURL)
	app.Get("/:slug", urlHandler.RedirectURL)
}
