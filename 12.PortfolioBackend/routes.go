package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app *fiber.App) {
	group := app.Group("/api/projects")

	group.Get("/", func(c *fiber.Ctx) error {
		projects := make([]Project, 0)

		if err := DB.Find(&projects).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Error getching projects: %v", err.Error()),
			})
		}

		return c.JSON(projects)
	})
	group.Post("/", func(c *fiber.Ctx) error {
		var project Project

		if err := c.BodyParser(&project); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Body is not valid!",
			})
		}

		if project.Title == "" || project.Description == "" || project.Language == "" || project.Link == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Some fields are missing!",
			})
		}

		if err := DB.Create(&project).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Error creating new project: %v", err.Error()),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Project created successfully!",
			"data":    project,
		})
	})
	group.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id", 0)

		if err != nil || id == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid document id!",
			})
		}

		var project Project
		if err := DB.Find(&project, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("Project with id %v not found!", id),
			})
		}

		if err := DB.Delete(&project).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to delete project!",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Project deleted successfully!",
			"data":    project,
		})
	})
	group.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id", 0)

		if err != nil || id == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid document id!",
			})
		}

		// Get Project with id
		var existing Project
		if err := DB.First(&existing, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("Project with id %v does not exist.", id),
			})
		}

		// Parse Body
		var input Project
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Body is not Valid!",
			})
		}
		input.ID = existing.ID

		if err := DB.Model(&existing).Updates(input).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update project.",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Project updated successfully!",
			"data":    existing,
		})
	})
}
