package profile

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID": user.ID,
		"Name":   user.Name,
		"Email":  user.Email,
		"Phone":  user.Phone,
		"Role":   user.Role,
	})
}
