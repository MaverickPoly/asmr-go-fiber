package main

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func GetAllQuotes(c *fiber.Ctx) error {
	return c.JSON(quotes)
}

func RandomQuote(c *fiber.Ctx) error {
	randomIndex := rand.Intn(len(quotes))
	return c.JSON(fiber.Map{
		"quote": quotes[randomIndex],
	})
}
