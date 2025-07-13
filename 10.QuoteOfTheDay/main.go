package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/quotes", GetAllQuotes)
	app.Get("/api/quotes/random", RandomQuote)

	fmt.Printf("App is running on port 8000")
	app.Listen(":8000")
}
