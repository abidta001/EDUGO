package routes

import (
	admin "edugo/controllers/Admin/courseManagement"
	tutormanagement "edugo/controllers/Admin/tutorManagement"
	"edugo/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminGroup := app.Group("/admin", middleware.JWTMiddleware, middleware.AdminMiddleware)

	adminGroup.Get("/requests", tutormanagement.ViewRequestTutor)
	adminGroup.Put("/verifytutor/:id", tutormanagement.VerifyTutor)
	adminGroup.Post("/category", admin.CreateCategory)
	adminGroup.Get("/view-category", admin.ViewCategoryAdmin)
}
