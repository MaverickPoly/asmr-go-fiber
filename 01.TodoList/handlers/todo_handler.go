package handlers

import (
	"fmt"
	"strconv"

	"01.TodoList/config"
	"01.TodoList/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos = make([]models.Todo, 0)
	query := `SELECT * FROM todos ORDER BY id`

	err := config.DB.Select(&todos, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	fmt.Println("Body: ", c.Body())

	query := `INSERT INTO todos (title) VALUES ($1) RETURNING id`
	err := config.DB.QueryRowx(query, todo.Title).Scan(&todo.ID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Failed to save todo: %s", err.Error())})
	}
	return c.Status(201).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	todoIdStr := c.Params("todoId")
	todoId, errInt := strconv.Atoi(todoIdStr)
	if errInt != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo Id!"})
	}

	query := `DELETE FROM todos WHERE id = $1 RETURNING id`
	err := config.DB.QueryRowx(query, todoId).Scan(&todoId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error deleting todo!"})
	}
	return c.JSON(fiber.Map{"message": "Todo deleted successfully!"})
}

func UpdateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	todoIdStr := c.Params("todoId")
	todoId, errInt := strconv.Atoi(todoIdStr)
	if errInt != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo Id!"})
	}

	query := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING id`
	var updatedId int
	err := config.DB.QueryRowx(query, todo.Title, todo.Completed, todoId).Scan(&updatedId)
	todo.ID = updatedId
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to update todo!"})
	}
	return c.JSON(todo)
}
