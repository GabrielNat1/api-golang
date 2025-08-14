package config

import (
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Path string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Path: getEnv("DB_PATH", "./database.db"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
