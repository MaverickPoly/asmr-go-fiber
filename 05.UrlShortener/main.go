package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	PORT := 3000

	app := fiber.New()
	ConnectDB()

	CreateRoutes(app)

	app.Listen(fmt.Sprintf(":%d", PORT))

	fmt.Printf("Listening on PORT %d\n", PORT)
}
