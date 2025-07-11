package main

import (
	"fmt"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UrlType struct {
	ID   int    `db:"id" json:"id "`
	Url  string `db:"url" json:"url"`
	Path string `db:"path" json:"path"`
}

// Shorten URL
func NewUrl(c *fiber.Ctx) error {
	fmt.Println("Entry NewUrl")
	body := new(UrlType)

	fmt.Println("Parsing Body...")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing body!",
		})
	}

	fmt.Println("Body Parsed")

	if !strings.Contains(body.Url, "http") && !strings.Contains(body.Url, "://") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL!",
		})
	}

	path := uuid.New()
	body.Path = path.String()

	fmt.Println("Path:", path)
	fmt.Println("Url:", body.Url)

	query := `INSERT INTO urls (url, path) VALUES (?, ?)`
	_, err := DB.Query(query, body.Url, body.Path)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(body)
}

// Get to URL From Shortened
func GetUrl(c *fiber.Ctx) error {
	path := c.Params("path")

	url := new(UrlType)
	query := `SELECT * FROM urls WHERE path = ? LIMIT 1`
	err := DB.Get(url, query, path)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not get url: %s", err.Error()),
		})
	}

	return c.Redirect(url.Url)
}
