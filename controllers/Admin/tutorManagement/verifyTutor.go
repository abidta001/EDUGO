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
	if err := config.DB.Where("user_id=?", userID).First(&userReq).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}
	userReq.Verified = true
	var user models.User
	if err := config.DB.Where("id=?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}
	user.Role = "tutor"

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered as Tutor!"})
}
