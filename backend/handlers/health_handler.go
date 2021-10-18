package handlers

import "github.com/gofiber/fiber/v2"

func APIStatusHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Status OK",
	})
}
