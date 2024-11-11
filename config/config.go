package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func LoadConfig() (*Config, error) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	log.Println("Env variables loaded from .env")

	cfg := &Config{
		Token: getEnv("DISCORD_BOT_TOKEN", ""),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
