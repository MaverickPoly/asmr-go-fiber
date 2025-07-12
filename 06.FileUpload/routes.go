package main

import "github.com/gofiber/fiber/v2"

func CreateRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/file/upload", UploadFile)
	api.Get("/file/get", GetFile)
}
