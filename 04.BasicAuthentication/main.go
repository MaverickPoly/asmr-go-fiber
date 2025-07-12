package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Constants
	PORT := os.Getenv("PORT")

	app := fiber.New()

	ConnectDB()
	CreateRoutes(app)

	log.Info(fmt.Sprintf("Server is running on Port %v\n", PORT))
	app.Listen(fmt.Sprintf(":%v", PORT))
}
