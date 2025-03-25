package authentication

import (
	"edugo/config"
	"edugo/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SignupUser(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	otp := GenerateOTP()
	if err := StoreOTP(input.Email, otp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store OTP"})
	}

	go func() {
		err := SendOTP(input.Email, otp)
		if err != nil {
			fmt.Println("Failed to send OTP", err)
		}
	}()

	password, err := HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: password,
		Role:     "student",
		Verified: false,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OTP sent to email. Verify to complete registration.",
	})
}
