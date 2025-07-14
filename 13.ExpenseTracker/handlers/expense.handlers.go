package handlers

import (
	"fmt"

	"13.ExpenseTracker/config"
	"13.ExpenseTracker/models"
	"github.com/gofiber/fiber/v2"
)

func FetchAllExpenses(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	expenses := make([]models.Expense, 0)

	if err := config.DB.Where("user_id = ?", userId).Find(&expenses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch expenses",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Fetched all expenses successfully!",
		"data":    expenses,
	})
}

func CreateExpense(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var expense models.Expense

	if err := c.BodyParser(&expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}

	if expense.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is missing!",
		})
	}

	expense.UserID = userId

	if err := config.DB.Create(&expense).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating expense: %v", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Expense created successfully!",
		"data":    expense,
	})
}

func DeleteExpense(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	expenseId, err := c.ParamsInt("expenseId", 0)

	if err != nil || expenseId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense id!",
		})
	}

	var expense models.Expense

	if err := config.DB.Where("user_id = ? AND id = ?", userId, expenseId).First(&expense).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Expense with id %v not found!", expenseId),
		})
	}

	if err := config.DB.Delete(&expense).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete expense!",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Expense with id %v deleted successfully!", expenseId),
		"data":    expense,
	})
}

func UpdateExpense(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	expenseId, err := c.ParamsInt("expenseId", 0)

	if err != nil || expenseId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expense id!",
		})
	}

	var dbExpense models.Expense

	if err := config.DB.Where("user_id = ? AND id = ?", userId, expenseId).First(&dbExpense).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Expense with id %v not found!", expenseId),
		})
	}

	var newExpense models.Expense
	newExpense.UserID = userId

	if err := c.BodyParser(&newExpense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}

	if err := config.DB.Model(&dbExpense).Updates(newExpense).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update expense!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Expense updated successfully!",
		"data":    newExpense,
	})
}
