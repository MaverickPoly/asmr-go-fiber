package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ConnectDB()
	app := fiber.New()

	SetupRoutes(app)

	fmt.Println("Server is running on PORT 3000")
	app.Listen(":3000")

}
