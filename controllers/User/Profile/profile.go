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

func EditUserProfile(c *fiber.Ctx) error {
	userId, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var input struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	var user models.User
	if err := config.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	user.Name = input.Name
	user.Phone = input.Phone

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update User"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Id":    user.ID,
		"Name":  user.Name,
		"Email": user.Email,
		"Phone": user.Phone,
	})
}
