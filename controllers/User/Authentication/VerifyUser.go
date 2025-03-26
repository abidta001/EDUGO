package authentication

import (
	"context"
	"edugo/config"
	"edugo/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func VerifyOTP(c *fiber.Ctx) error {
	var input struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx := context.Background()
	key := "otp:" + input.Email
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

	if err := config.DB.Model(&models.User{}).Where("email = ?", input.Email).Update("verified", true).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify user"})
	}

	redisClient.Del(ctx, key)

	go SendMail(input.Email)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User verified successfully. You can now log in."})

}
