package courses

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func ViewCourses(c *fiber.Ctx) error {

	var courses []models.Course

	if err := config.DB.Preload("Category").Preload("Tutor.User").Find(&courses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch course"})
	}
	var response []fiber.Map
	for _, course := range courses {
		response = append(response, fiber.Map{
			"ID":          course.ID,
			"Title":       course.Title,
			"Description": course.Description,
			"Price":       course.Price,
			"Category":    course.Category.Name,
			"Tutor":       course.Tutor.User.Name,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
