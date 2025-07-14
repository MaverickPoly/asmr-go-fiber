package main

import (
	"14.FileFolderExplorer/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()

	api := app.Group("/api")

	api.Post("/items", createItem)
	api.Get("/items/:parentId?", getItemsByParent) // ? means optional
	api.Delete("/items/:id", deleteItem)

	app.Listen(":3000")
}
