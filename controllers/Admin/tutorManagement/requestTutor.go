package tutormanagement

import (
	"edugo/config"
	"edugo/models"

	"github.com/gofiber/fiber/v2"
)

func ViewRequestTutor(c *fiber.Ctx) error {
	var tutorRequests []models.Tutor

	if err := config.DB.Where("verified = ?", false).Find(&tutorRequests).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tutor requests",
		})
	}

	type TutorRequestResponse struct {
		ID             uint   `json:"id"`
		Qualifications string `json:"qualifications"`
		Expertise      string `json:"expertise"`
		Bio            string `json:"bio"`
		Experience     int    `json:"experience"`
		Availability   string `json:"availability"`
	}

	var response []TutorRequestResponse
	for _, tutor := range tutorRequests {
		response = append(response, TutorRequestResponse{
			ID:             tutor.ID,
			Qualifications: tutor.Qualifications,
			Expertise:      tutor.Expertise,
			Bio:            tutor.Bio,
			Experience:     tutor.Experience,
			Availability:   tutor.Availability,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"requests": response,
	})
}
