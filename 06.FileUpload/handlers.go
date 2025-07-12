package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	formFile, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get uploaded file!",
		})
	}

	file := File{
		Filename: formFile.Filename,
		Size:     float64(formFile.Size),
		Path:     fmt.Sprintf("./uploads/%s", formFile.Filename),
	}

	if err := c.SaveFile(formFile, file.Path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error saving file: %s", err.Error()),
		})
	}

	if result := DB.Create(&file); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to save file in db: %s", result.Error.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "File created successfully!",
	})
}

func GetFile(c *fiber.Ctx) error {
	filename := strings.Trim(c.Query("filename", ""), "\"")

	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid filename!",
		})
	}

	fmt.Println(filename)

	file := new(File)

	result := DB.Where("filename = ?", filename).First(&file)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("File with name %s does not exists!", filename),
		})
	}

	return c.Download(file.Path)
}
