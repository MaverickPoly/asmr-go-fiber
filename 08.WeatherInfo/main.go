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

	PORT := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", IndexHandler)
	app.Get("/weather/:city", CityWeather)

	log.Info("Server is listening on PORT", PORT)
	app.Listen(fmt.Sprintf(":%v", PORT))
}
