package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mofodox/sena/database"
	"github.com/mofodox/sena/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("unable to load the .env file: %s\n", err.Error())
	} else {
		log.Println("Successfully load the .env file")
	}

	// Initialise the connection to the database
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("unable to connect to the database")
	}

	app := fiber.New()

	// middlewares
	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
