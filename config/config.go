package config

import "os"

type Config struct {
	PORT       string
	DB_DSN     string
	JWT_SECRET string
}

func Load() *Config {

	return &Config{
		DB_DSN:     getEnv("DB_DSN", ""),
		JWT_SECRET: getEnv("JWT_SECRET", ""),
		PORT:       getEnv("PORT", "9000"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
