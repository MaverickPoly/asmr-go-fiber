package main

import (
	"02.NotesApp/config"
	"02.NotesApp/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
