package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	// Constants
	godotenv.Load()

	PORT := os.Getenv("PORT")

	ConnectDB()
	app := fiber.New()
	SetupRoutes(app)

	log.Info(fmt.Sprintf("Server is up and running on PORT %v", PORT))
	app.Listen(fmt.Sprintf(":%v", PORT))

}
