package config

import "os"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	App      AppConfig
	LogLevel string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URI  string
	Name string
}

type AppConfig struct {
	BaseURL        string
	AllowedOrigins string
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Database: DatabaseConfig{
			URI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Name: getEnv("MONGODB_NAME", "db"),
		},
		App: AppConfig{
			BaseURL:        getEnv("BASE_URL", "http://localhost:8080"),
			AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
		},
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
