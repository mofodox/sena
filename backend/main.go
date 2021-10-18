package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mofodox/sena/database"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("unable to load the .env file: %s\n", err.Error())
	} else {
		log.Println("Successfully load the .env file")
	}

	// Initialise the connection to the database
	database.ConnectDB()

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
