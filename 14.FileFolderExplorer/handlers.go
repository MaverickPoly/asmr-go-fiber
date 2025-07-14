package main

import (
	"14.FileFolderExplorer/database"
	"14.FileFolderExplorer/models"
	"github.com/gofiber/fiber/v2"
)

func createItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	if item.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
	}

	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create item"})
	}

	return c.Status(201).JSON(item)
}

func getItemsByParent(c *fiber.Ctx) error {
	parentID := c.Params("parentId")
	var items []models.Item

	if parentID == "" {
		database.DB.Where("parent_id IS NULL").Find(&items)
	} else {
		database.DB.Where("parent_id = ?", parentID).Find(&items)
	}

	return c.JSON(items)
}

func deleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	if item.IsFolder {
		database.DB.Where("parent_id = ?", item.ID).Delete(&models.Item{})
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete item"})
	}

	return c.JSON(fiber.Map{"message": "Item deleted", "item": item})
}
