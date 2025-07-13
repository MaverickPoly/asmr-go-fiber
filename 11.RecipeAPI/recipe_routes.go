package main

import "github.com/gofiber/fiber/v2"

func RecipeRoutes(api fiber.Router) {
	recipeRoutes := api.Group("/recipes")

	// Endpoints
	recipeRoutes.Get("/", GetAllRecipes)
	recipeRoutes.Post("/", CreateRecipe)
	recipeRoutes.Get("/:recipeId", GetRecipe)
	recipeRoutes.Put("/:recipeId", UpdateRecipe)
	recipeRoutes.Delete("/:recipeId", DeleteRecipe)
}
