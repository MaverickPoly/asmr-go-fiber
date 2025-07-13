package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllRecipes(c *fiber.Ctx) error {
	nameQuery := c.Query("name", "")
	ingredientQuery := c.Query("ingridient", "")

	query := DB.Model(&Recipe{})

	if nameQuery != "" {
		query = query.Where("name ILIKE ?", "%"+nameQuery+"%")
	}
	if ingredientQuery != "" {
		query = query.Where("ingredients ILIKE ?", "%"+ingredientQuery+"%")
	}

	var recipes []Recipe

	if err := query.Find(&recipes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch recipes!",
		})
	}

	return c.JSON(recipes)
}

func CreateRecipe(c *fiber.Ctx) error {
	var recipe Recipe

	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}

	if recipe.Name == "" || recipe.Description == "" || recipe.Ingredients == "" || recipe.Steps == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Some fields are missing!",
		})
	}

	if err := DB.Create(&recipe).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating recipe!",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Recipe created successfully!",
		"data":    recipe,
	})
}

func GetRecipe(c *fiber.Ctx) error {
	recipeId, err := c.ParamsInt("recipeId", 0)

	if err != nil || recipeId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid recipe id!",
		})
	}

	var recipe Recipe
	if err := DB.First(&recipe, recipeId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Recipe with id %v does not exist!", recipeId),
		})
	}

	return c.JSON(recipe)
}

func UpdateRecipe(c *fiber.Ctx) error {
	recipeId, err := c.ParamsInt("recipeId", 0)
	if err != nil || recipeId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid recipe id!",
		})
	}

	var recipe Recipe
	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body!",
		})
	}
	recipe.ID = uint(recipeId)

	if err := DB.Model(&recipe).Updates(recipe).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating recipe!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Recipe updated successfully!",
		"data":    recipe,
	})
}

func DeleteRecipe(c *fiber.Ctx) error {
	recipeId, err := c.ParamsInt("recipeId", 0)
	if err != nil || recipeId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid recipe id!",
		})
	}

	var recipe Recipe
	if err := DB.First(&recipe, recipeId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Recipe with id %v not found!", recipeId),
		})
	}

	if err := DB.Delete(&Recipe{}, recipeId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete recipe!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Deleted recipe successfully!",
		"data":    recipe,
	})

}
