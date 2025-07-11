package main

import "github.com/gofiber/fiber/v2"

func CreateRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/new", NewUrl)
	api.Get("/:path", GetUrl)
}
