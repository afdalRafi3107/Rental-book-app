package main

import (
	"os"
	"rental-book/internal/config"
	"rental-book/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()
	routes.BookRoute(r)
	config.ConnectDB()
	config.EnvLoad()
	r.Listen(":"+os.Getenv("PORT"))
}