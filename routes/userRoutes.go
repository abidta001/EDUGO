package routes

import (
	authentication "edugo/controllers/User/Authentication"
	courses "edugo/controllers/User/Courses"
	profile "edugo/controllers/User/Profile"
	"edugo/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/signup", authentication.SignupUser)
	auth.Post("/login", authentication.LoginUser)
	auth.Post("/verify", authentication.VerifyOTP)
	auth.Post("/resend", authentication.ResendOTP)
}

func ProfileRoutes(app *fiber.App) {
	profileGroup := app.Group("/profile", middleware.JWTMiddleware)

	profileGroup.Get("/view", profile.GetUserProfile)
	profileGroup.Put("/edit", profile.EditUserProfile)
	profileGroup.Post("/reset-password", profile.ResetPasswordOTP)
	profileGroup.Post("/change-password", profile.ChangePassword)
}

func NormalRoutes(app *fiber.App) {
	app.Get("/category", courses.ViewCategory)
	app.Get("/courses", courses.ViewCourses)

}
