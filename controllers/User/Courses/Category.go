package courses

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func ViewCategory(c *fiber.Ctx) error {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error finding categories"})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var category models.Category
	category.Name = input.Name
	category.Description = input.Description

	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error creating category"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"CategoryID":  category.ID,
		"Name":        category.Name,
		"Description": category.Description,
	})
}
