package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mofodox/sena/handlers"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/health", handlers.APIStatusHandler)
	apiV1.Get("/todos", handlers.GetTodos)
	apiV1.Post("/todos", handlers.CreateTodo)
}
