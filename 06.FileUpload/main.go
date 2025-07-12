package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	InitDB()
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	os.Mkdir("uploads", 0755)

	CreateRoutes(app)

	log.Info("App is listening on PORT 3000")
	app.Listen(":3000")
}
