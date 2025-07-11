package router

import (
	"03.BookCatalog/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(api fiber.Router) {
	bookRoutes := api.Group("/books")
	bookRoutes.Get("/", handlers.GetBooks)
	bookRoutes.Post("/", handlers.CreateBook)
	bookRoutes.Get("/:bookId", handlers.GetBook)
	bookRoutes.Delete("/:bookId", handlers.DeleteBook)
	bookRoutes.Put("/:bookId", handlers.UpdateBook)
}
