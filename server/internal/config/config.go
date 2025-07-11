package config

import (
	"time"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	App       AppConfig
	RateLimit RateLimitConfig
	LogLevel  string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URI  string
	Name string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type AppConfig struct {
	AllowedOrigins string
}

type RateLimitConfig struct {
	Enabled   bool
	Requests  int
	Window    time.Duration
	BlockTime time.Duration
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "1209"),
		},
		Database: DatabaseConfig{
			URI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Name: getEnv("MONGODB_NAME", "db"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		App: AppConfig{
			AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
		},
		RateLimit: RateLimitConfig{
			Enabled:   getEnvAsBool("RATE_LIMIT_ENABLED", true),
			Requests:  getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
			Window:    time.Duration(getEnvAsInt("RATE_LIMIT_WINDOW_MINUTES", 1)) * time.Minute,
			BlockTime: time.Duration(getEnvAsInt("RATE_LIMIT_BLOCK_MINUTES", 5)) * time.Minute,
		},
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}, nil
}
