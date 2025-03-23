package main

import (
	"edugo/config"
	authentication "edugo/controllers/User/Authentication"
	profile "edugo/controllers/User/Profile"
	"edugo/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.InitDB()

	//Authentication
	app.Post("/signup", authentication.SignupUser)
	app.Post("/login", authentication.LoginUser)

	//Profile
	app.Get("/profile/view", middleware.JWTMiddleware, profile.GetUserProfile)

	log.Fatal(app.Listen(":3000"))
}
