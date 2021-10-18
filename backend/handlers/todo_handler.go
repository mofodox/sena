package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mofodox/sena/database"
	"github.com/mofodox/sena/models"
)

// GetTodos retrieve all todos
func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	database.DB.Model(&models.Todo{}).Find(&todos)

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "ok",
		"todos":   todos,
	})
}

// CreateTodo insert todo
func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// set the todo ID using UUID
	todo.ID = uuid.New()

	database.DB.Create(&todo)

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "ok",
		"todo":    todo,
	})
}
