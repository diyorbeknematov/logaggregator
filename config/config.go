package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServerAddress string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	var config Config

	config.ServerAddress = cast.ToString(coalesce("SERVER_ADDRESS", ":5000"))

	return config
}

func coalesce(key string, defaults interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaults
}
