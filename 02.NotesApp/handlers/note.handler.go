package handlers

import (
	"fmt"
	"strconv"

	"02.NotesApp/config"
	"02.NotesApp/models"
	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	var notes = make([]models.Note, 0)
	query := `SELECT * FROM notes ORDER BY id`

	err := config.DB.Select(&notes, query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(notes)
}

func CreateNote(c *fiber.Ctx) error {
	note := new(models.Note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse body"})
	}
	fmt.Println("Note:", note)

	query := `INSERT INTO notes (title) VALUES (?)`
	res, err := config.DB.Exec(query, note.Title)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to insert note!",
		})
	}

	id, _ := res.LastInsertId()
	note.ID = int(id)

	return c.Status(201).JSON(note)
}

func GetNote(c *fiber.Ctx) error {
	noteId, errInt := strconv.Atoi(c.Params("noteId"))
	if errInt != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Note ID",
		})
	}

	note := new(models.Note)

	query := `SELECT * FROM notes WHERE id = ? LIMIT 1`
	err := config.DB.Get(note, query, noteId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Error getting note: %s", err.Error())})
	}

	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	noteId, err := strconv.Atoi(c.Params("noteId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Note ID",
		})
	}

	query := `DELETE FROM notes WHERE id = ?`
	if _, err := config.DB.Exec(query, noteId); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete note!",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Note with ID %d deleted successfully!", noteId),
	})
}

func UpdateNote(c *fiber.Ctx) error {
	noteId, err := strconv.Atoi(c.Params("noteId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	note := new(models.Note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	query := `UPDATE notes SET title = ? WHERE id = ?`
	if _, err := config.DB.Exec(query, note.Title, noteId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error updating note: %s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Note with id %d was updated successfully!", noteId),
	})
}
