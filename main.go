package main

import (
	"edugo/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.InitDB()

	log.Fatal(app.Listen(":3000"))
}
