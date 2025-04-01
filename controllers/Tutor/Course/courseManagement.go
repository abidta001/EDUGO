package course

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func CreateCourse(c *fiber.Ctx) error {
	var course models.Course

	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}
	userID, err := c.Locals("userID").(uint)
	if !err {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	course.TutorID = userID
	if err := config.DB.Create(&course).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create course"})
	}
	var savedCourse models.Course
	if err := config.DB.Preload("Category").Preload("Tutor.User").First(&savedCourse, course.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve course details"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ID":          savedCourse.ID,
		"Title":       savedCourse.Title,
		"Description": savedCourse.Description,
		"Price":       savedCourse.Price,
		"Category":    savedCourse.Category.Name,
		"Tutor":       savedCourse.Tutor.User.Name,
	})
}
