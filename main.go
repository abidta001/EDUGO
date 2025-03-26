package main

import (
	"edugo/config"
	admin "edugo/controllers/Admin/courseManagement"
	tutormanagement "edugo/controllers/Admin/tutorManagement"
	tutor "edugo/controllers/Tutor"
	profiletutor "edugo/controllers/Tutor/ProfileTutor"
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
	app.Post("/resend", authentication.ResendOTP)

	//Profile
	app.Get("/profile/view", middleware.JWTMiddleware, profile.GetUserProfile)
	app.Put("profile/edit", middleware.JWTMiddleware, profile.EditUserProfile)

	//Course
	app.Get("/category", courses.ViewCategory)
	//Give this to admin

	//Tutor
	app.Post("/request", middleware.JWTMiddleware, tutor.RequestTutor)
	app.Get("/viewtutor", middleware.JWTMiddleware, profiletutor.ViewTutorProfile)
	//admin

	//Admin
	app.Get("/request", middleware.JWTMiddleware, middleware.AdminMiddleware, tutormanagement.ViewRequestTutor)
	app.Put("/verifytutor/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, tutormanagement.VerifyTutor)
	app.Post("/category", middleware.JWTMiddleware, middleware.AdminMiddleware, admin.CreateCategory)
	app.Get("/view-category", middleware.JWTMiddleware, middleware.AdminMiddleware, admin.ViewCategoryAdmin)

	log.Fatal(app.Listen(":3000"))
}
