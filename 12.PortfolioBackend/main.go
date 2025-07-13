package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	ConnectDB()
	app := fiber.New()

	ProjectRoutes(app)

	log.Info("Server is listening on port 3000")
	app.Listen(":3000")
}
