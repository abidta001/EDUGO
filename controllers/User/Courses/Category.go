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
