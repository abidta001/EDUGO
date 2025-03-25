package tutorManagement

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func VerifyTutor(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var userReq models.Tutor
	if err := config.DB.Where("user_id = ?", userID).First(&userReq).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tutor request not found"})
	}

	userReq.Verified = true
	if err := config.DB.Save(&userReq).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify tutor request"})
	}

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.Role = "tutor"
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user role"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered as Tutor!"})
}
