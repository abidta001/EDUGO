package profile

import (
	"context"
	"edugo/config"
	authentication "edugo/controllers/User/Authentication"
	"edugo/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResetPasswordOTP(c *fiber.Ctx) error {
	var user models.User
	userIDInterface := c.Locals("userID")
	userID, ok := userIDInterface.(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if err := config.DB.Where("id=?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	otp := authentication.GenerateOTP()
	if err := authentication.SendOTP(user.Email, otp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send OTP"})
	}

	if err := authentication.StoreOTP(user.Email, otp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store OTP"})
	}

	fmt.Println("OTP :", otp)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP sent successfully!"})
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func ChangePassword(c *fiber.Ctx) error {
	var input struct {
		OTP             string `json:"otp"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if input.Password != input.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Passwords do not match"})
	}

	userIDInterface := c.Locals("userID")
	userID, ok := userIDInterface.(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := config.DB.Where("id=?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user"})
	}

	ctx := context.Background()
	key := "otp:" + user.Email
	storedOTP, err := redisClient.Get(ctx, key).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "OTP expired or invalid"})
	} else if err != nil {
		fmt.Println("Error fetching OTP:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	if storedOTP != input.OTP {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	hashpassword, err := authentication.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user.Password = hashpassword
	config.DB.Save(&user)

	redisClient.Del(ctx, key)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password changed successfully"})
}
