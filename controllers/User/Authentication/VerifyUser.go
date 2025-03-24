package authentication

import (
	"context"
	"edugo/config"
	"edugo/models"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOTP(email, otp string) error {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal("Error loading .env file")
	}
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: EduGo OTP Verification\n\nYour OTP is: " + otp)
	fmt.Println("Message", string(message))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}

func StoreOTP(email, otp string) error {
	ctx := context.Background()
	key := "otp:" + email
	err := redisClient.Set(ctx, key, otp, 5*time.Minute).Err()
	return err
}

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User verified successfully. You can now log in."})
}
