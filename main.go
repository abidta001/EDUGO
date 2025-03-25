package main

import (
	"edugo/config"
	tutormanagement "edugo/controllers/Admin/tutorManagement"
	tutor "edugo/controllers/Tutor"
	authentication "edugo/controllers/User/Authentication"
	courses "edugo/controllers/User/Courses"
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
	app.Post("/verify", authentication.VerifyOTP)

	//Profile
	app.Get("/profile/view", middleware.JWTMiddleware, profile.GetUserProfile)
	app.Put("profile/edit", middleware.JWTMiddleware, profile.EditUserProfile)

	//Course
	app.Get("/category", courses.ViewCategory)
	app.Post("/category", middleware.JWTMiddleware, middleware.AdminMiddleware, courses.CreateCategory) //Give this to admin

	//Tutor
	app.Post("/request", middleware.JWTMiddleware, tutor.RequestTutor)
	app.Get("/request", middleware.JWTMiddleware, middleware.AdminMiddleware, tutormanagement.ViewRequestTutor) //admin

	//Admin
	app.Put("/verifytutor/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, tutormanagement.VerifyTutor)

	log.Fatal(app.Listen(":3000"))
}
