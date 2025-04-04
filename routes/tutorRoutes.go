package routes

import (
	tutor "edugo/controllers/Tutor"
	course "edugo/controllers/Tutor/Course"
	profiletutor "edugo/controllers/Tutor/ProfileTutor"
	"edugo/middleware"

	"github.com/gofiber/fiber/v2"
)

func TutorRoutes(app *fiber.App) {
	tutorGroup := app.Group("/tutor", middleware.JWTMiddleware)

	tutorGroup.Post("/request", tutor.RequestTutor)
	tutorGroup.Get("/view", profiletutor.ViewTutorProfile)
	tutorGroup.Post("/course", course.CreateCourse)
}
