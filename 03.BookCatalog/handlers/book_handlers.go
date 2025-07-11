package handlers

import (
	"fmt"

	"03.BookCatalog/config"
	"03.BookCatalog/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books := make([]models.Book, 0)
	rating := c.QueryInt("rating", -1)
	var query string
	var err error

	if rating != -1 {
		query = `SELECT * FROM books WHERE rating = $1 ORDER BY id`
		err = config.DB.Select(&books, query, rating)
	} else {
		query = `SELECT * FROM books ORDER BY id`
		err = config.DB.Select(&books, query)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error getting books: %s", err.Error()),
		})
	}

	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing body!",
		})
	}

	query := `
		INSERT INTO books (title, author, description, rating) VALUES ($1, $2, $3, $4) RETURNING id
	`
	err := config.DB.QueryRowx(query, book.Title, book.Author, book.Description, book.Rating).Scan(&book.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert book!",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	bookId, errInt := c.ParamsInt("bookId")

	if errInt != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	book := new(models.Book)

	query := `SELECT * FROM books WHERE id = $1 LIMIT 1`
	err := config.DB.Get(book, query, bookId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not get book with id %d", bookId),
		})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	bookId, errInt := c.ParamsInt("bookId")

	if errInt != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var book models.Book

	query := `DELETE FROM books WHERE id = $1 RETURNING *`

	err := config.DB.Get(&book, query, bookId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error deleting book with id %d", bookId),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Book with ID %d deleted successfully!", bookId),
		"data":    book,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	bookId, errInt := c.ParamsInt("bookId")

	if errInt != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse body!",
		})
	}

	query := `UPDATE books SET title = $1, author = $2, description = $3, rating = $4 WHERE id = $5 RETURNING id`
	err := config.DB.QueryRowx(query, book.Title, book.Author, book.Description, book.Rating, bookId).Scan(&book.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error updaing book with id %d", bookId),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Book with ID %d updated successfully!", bookId),
		"data":    book,
	})
}
