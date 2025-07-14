package main

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/lareii/siker.im/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	logger *zap.Logger
}

func NewServer(app *fiber.App, config *config.Config, logger *zap.Logger) *Server {
	return &Server{
		app:    app,
		config: config,
		logger: logger,
	}
}

func (s *Server) Start() {
	go func() {
		s.logger.Info("Server starting on port " + s.config.Server.Port)
		if err := s.app.Listen(":" + s.config.Server.Port); err != nil {
			s.logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
