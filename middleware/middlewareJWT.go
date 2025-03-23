package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 || tokenString[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["id"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	c.Locals("userID", uint(claims["id"].(float64)))

	return c.Next()
}
