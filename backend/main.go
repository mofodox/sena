package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// middlewares
	app.Use(logger.New())

	app.Get("/health", apiStatusHandler)

	app.Listen(":8080")
}

func apiStatusHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Status OK",
	})
}
