package routes

import (
	"13.ExpenseTracker/handlers"
	"13.ExpenseTracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	router := api.Group("/auth")

	router.Post("/login", handlers.HandleLogin)
	router.Post("/register", handlers.HandleRegister)
	router.Post("/logout", handlers.HandleLogout)
	router.Get("/me", middleware.AuthRequired, handlers.GetMe)
}
