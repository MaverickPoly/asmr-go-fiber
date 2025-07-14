package routes

import (
	"13.ExpenseTracker/handlers"
	"13.ExpenseTracker/middleware"
	"github.com/gofiber/fiber/v2"
)

func ExpenseRoutes(api fiber.Router) {
	router := api.Group("/expenses")

	router.Get("/", middleware.AuthRequired, handlers.FetchAllExpenses)
	router.Post("/", middleware.AuthRequired, handlers.CreateExpense)
	router.Delete("/:expenseId", middleware.AuthRequired, handlers.DeleteExpense)
	router.Put("/:expenseId", middleware.AuthRequired, handlers.UpdateExpense)
}
