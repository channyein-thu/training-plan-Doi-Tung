package router

import (
	"training-plan-api/container"

	"github.com/gofiber/fiber/v2"
)

func StaffRoutes(r fiber.Router, deps *container.AppDependencies) {

	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	

	// // Records (own)
	r.Get("/records", deps.RecordController.FindByCurrentUser)
	r.Get("/records/:id", deps.RecordController.FindById)

	// // Certificates
	r.Get("/certificates", deps.CertificateController.FindByCurrentUser) // only approved certificates
	r.Post("/certificates", deps.CertificateController.Upload)
	r.Delete("/certificates/:id", deps.CertificateController.Delete)
}
