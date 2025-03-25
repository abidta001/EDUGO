package tutor

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func RequestTutor(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var input struct {
		Qualifications string `json:"qualifications"`
		Expertise      string `json:"expertise"`
		Bio            string `json:"bio"`
		Experience     int    `json:"experience"`
		Availability   string `json:"availability"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	tutor := models.Tutor{
		UserID:         userID,
		Qualifications: input.Qualifications,
		Expertise:      input.Expertise,
		Bio:            input.Bio,
		Experience:     input.Experience,
		Availability:   input.Availability,
		Verified:       false,
	}
	if err := config.DB.Create(&tutor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Request Pending"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"error": "Request send! waiting for admin approval"})
}
