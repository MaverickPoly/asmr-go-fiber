package main

import (
	"03.BookCatalog/config"
	"03.BookCatalog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.ConnectDB()
	app := fiber.New()      // Routes
	router.SetupRoutes(app) // Setup Routes
	app.Listen(":8000")
	log.Info("App is running on port 3000!")
}
