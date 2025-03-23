package authentication

import (
	"edugo/config"
	"edugo/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SignupUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request!"})
	}

	if err := user.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Password hashing failed"})
	}
	user.Password = hashedPassword
	user.Role = "student"

	if err := config.DB.Create(&user).Error; err != nil {
		fmt.Println("Error creating user in DB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	otp := GenerateOTP()
	if err := StoreOTP(fmt.Sprintf("%d", user.ID), otp); err != nil {
		fmt.Println("Error storing OTP:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to store OTP"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"UserID":  user.ID,
		"Name":    user.Name,
		"Email":   user.Email,
		"Phone":   user.Phone,
		"Role":    user.Role,
		"Otp":     otp,
		"Message": "Verify user for login!",
	})
}
