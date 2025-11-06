package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	DB        DBConfig
	Server    ServerConfig
	JWTSecret string
	Env       string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DB: DBConfig{
			DSN: getEnv("DB_URL", "postgres://postgres:Tsiva22@@localhost:5432/productdb?sslmode=disable"),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "3001"),
		},
		JWTSecret: getEnv("JWT_SECRET", "fallback-secret"),
		Env:       getEnv("ENV", "development"),
	}

	log.Printf("✅ Loaded config for %s environment", cfg.Env)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		log.Printf("✅ Loaded %s from environment", key)
		return value
	}
	return fallback
}
