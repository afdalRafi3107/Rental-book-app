package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	DBUrl string
}

func LoadEnv() *Config{
	err:= godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Port: os.Getenv("PORT"),
		DBUrl:os.Getenv("DB_URL"),
	}
}
func getEnv (key, fallback string) string{
		if value,  exists := os.LookupEnv(key); exists{
			return  value
		}
		return  fallback
	}