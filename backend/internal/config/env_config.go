package config

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env not found", err)
	}
}