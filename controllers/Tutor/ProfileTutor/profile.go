package profiletutor

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func ViewTutorProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var tutor models.Tutor
	if err := config.DB.Preload("User").Where("user_id = ?", userID).First(&tutor).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tutor profile not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID":         tutor.User.ID,
		"Name":           tutor.User.Name,
		"Email":          tutor.User.Email,
		"Phone":          tutor.User.Phone,
		"Qualifications": tutor.Qualifications,
		"Expertise":      tutor.Expertise,
		"Bio":            tutor.Bio,
		"Experience":     tutor.Experience,
		"Availability":   tutor.Availability,
	})
}
