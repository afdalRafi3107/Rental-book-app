package http

import (
	"RentalBook/internal/config"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberApp(cfg *config.Config) *fiber.App{
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	return app
}