package main

import "github.com/gofiber/fiber/v2"

func CreateRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", HandleRegister)
	api.Post("/auth/login", HandleLogin)
	api.Post("/auth/logout", AuthRequired, HandleLogout)
	api.Get("/auth/me", AuthRequired, HandleGetUser)
}
