package config

import (
	"strings"
	"time"

	"github.com/lareii/siker.im/internal/utils"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	App       AppConfig
	RateLimit RateLimitConfig
	Turnstile string
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
	AllowedOrigins []string
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
			Port: utils.GetEnv("PORT", "1209"),
		},
		Database: DatabaseConfig{
			URI:  utils.GetEnv("MONGODB_URI", "mongodb://mongodb:27017"),
			Name: utils.GetEnv("MONGODB_NAME", "db"),
		},
		Redis: RedisConfig{
			Host:     utils.GetEnv("REDIS_HOST", "redis"),
			Port:     utils.GetEnv("REDIS_PORT", "6379"),
			Password: utils.GetEnv("REDIS_PASSWORD", ""),
			DB:       utils.GetEnvAsInt("REDIS_DB", 0),
		},
		App: AppConfig{
			AllowedOrigins: strings.Split(utils.GetEnv("ALLOWED_ORIGINS", "*"), ","),
		},
		RateLimit: RateLimitConfig{
			Enabled:   utils.GetEnvAsBool("RATE_LIMIT_ENABLED", true),
			Requests:  utils.GetEnvAsInt("RATE_LIMIT_REQUESTS", 100),
			Window:    time.Duration(utils.GetEnvAsInt("RATE_LIMIT_WINDOW_MINUTES", 1)) * time.Minute,
			BlockTime: time.Duration(utils.GetEnvAsInt("RATE_LIMIT_BLOCK_MINUTES", 5)) * time.Minute,
		},
		Turnstile: utils.GetEnv("TURNSTILE_SECRET", "1x0000000000000000000000000000000AA"),
		LogLevel:  utils.GetEnv("LOG_LEVEL", "info"),
	}, nil
}
