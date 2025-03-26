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
)

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

func SendMail(email string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	subject := "Subject: Welcome to EduGo!\n"
	body := "Hello,\n\nCongratulations! You have successfully registered on EduGo.\n\n" +
		"We are excited to have you on board. Start exploring courses and enhance your learning journey!\n\n" +
		"Best regards,\nEduGo Team"

	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}
	return nil
}

func ResendOTP(c *fiber.Ctx) error {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	otp := GenerateOTP()

	if err := SendOTP(input.Email, otp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send OTP email"})
	}

	if err := StoreOTP(input.Email, otp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store OTP"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP resent successfully"})
}
