package main

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/lareii/siker.im/internal/config"
	"github.com/lareii/siker.im/internal/database"
	"github.com/lareii/siker.im/internal/handlers"
	"github.com/lareii/siker.im/internal/middleware"
	"github.com/lareii/siker.im/internal/repository"
	"github.com/lareii/siker.im/internal/services"
	"go.uber.org/zap"
)

type dependencies struct {
	DB         *database.MongoDB
	Redis      *database.Redis
	URLHandler *handlers.URLHandler
	Middleware *MiddlewareContainer
}

type MiddlewareContainer struct {
	RateLimiter *middleware.RateLimiter
	Turnstile   *middleware.TurnstileMiddleware
}

func initApp(cfg *config.Config, logger *zap.Logger) (*fiber.App, func(), error) {
	deps, cleanup, err := initDeps(cfg, logger)
	if err != nil {
		return nil, cleanup, err
	}

	app := createFiberApp(logger)
	setupRoutes(app, deps, cfg)

	return app, cleanup, nil
}

func initDeps(cfg *config.Config, logger *zap.Logger) (*dependencies, func(), error) {
	db, err := initDB(cfg, logger)
	if err != nil {
		return nil, nil, err
	}

	redis, err := initRedis(cfg, logger)
	if err != nil {
		return nil, func() { db.Disconnect(context.Background()) }, err
	}

	urlRepo := repository.NewURLRepository(db)
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService, logger)
	middlewareContainer := &MiddlewareContainer{
		RateLimiter: middleware.NewRateLimiter(redis, &cfg.RateLimit, logger),
		Turnstile:   middleware.NewTurnstileMiddleware(cfg.Turnstile),
	}

	deps := &dependencies{
		DB:         db,
		Redis:      redis,
		URLHandler: urlHandler,
		Middleware: middlewareContainer,
	}

	cleanup := func() {
		redis.Close()
		db.Disconnect(context.Background())
	}

	return deps, cleanup, nil
}

func initDB(cfg *config.Config, logger *zap.Logger) (*database.MongoDB, error) {
	db, err := database.NewMongoDB(cfg.Database.URI, cfg.Database.Name)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB", zap.String("uri", cfg.Database.URI))
	return db, nil
}

func initRedis(cfg *config.Config, logger *zap.Logger) (*database.Redis, error) {
	redis, err := database.NewRedis(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to Redis", zap.String("host", cfg.Redis.Host), zap.String("port", cfg.Redis.Port))
	return redis, nil
}

func createFiberApp(logger *zap.Logger) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:      "siker.im",
		ServerHeader: "Fiber",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			logger.Error("Request error", zap.Error(err))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})
}
