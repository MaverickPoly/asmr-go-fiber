package main

import (
	"fmt"
	"os"

	"13.ExpenseTracker/config"
	"13.ExpenseTracker/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Constants
	PORT := os.Getenv("PORT")

	config.ConnectDB()
	app := fiber.New()
	routes.SetupRoutes(app)

	log.Info("Server is running on PORT ", PORT)
	app.Listen(fmt.Sprintf(":%v", PORT))
}
