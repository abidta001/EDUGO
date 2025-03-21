package main

import (
	"edugo/config"
	authentication "edugo/controllers/Authentication"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.InitDB()

	//Authentication
	app.Post("/signup", authentication.SignupUser)
	app.Post("/login", authentication.LoginUser)

	log.Fatal(app.Listen(":3000"))
}
