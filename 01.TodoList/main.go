package main

import (
	"01.TodoList/config"
	"01.TodoList/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
