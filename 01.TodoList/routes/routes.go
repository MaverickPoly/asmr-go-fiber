package routes

import (
	"01.TodoList/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Todos
	api.Get("/todos", handlers.GetTodos)
	api.Post("/todos", handlers.CreateTodo)
	api.Delete("/todos/:todoId", handlers.DeleteTodo)
	api.Put("/todos/:todoId", handlers.UpdateTodo)
}
