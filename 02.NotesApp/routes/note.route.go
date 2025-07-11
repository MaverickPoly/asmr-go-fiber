package routes

import (
	"02.NotesApp/handlers"
	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(api fiber.Router) fiber.Router {
	router := api.Group("/notes")

	// Notes
	router.Get("/", handlers.GetNotes)
	router.Post("/", handlers.CreateNote)
	router.Get("/:noteId", handlers.GetNote)
	router.Delete("/:noteId", handlers.DeleteNote)
	router.Put("/:noteId", handlers.UpdateNote)

	return router
}
