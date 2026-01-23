package router

import (
	"training-plan-api/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router, authController *controller.AuthController) {
	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	// r.Post("/auth/login", authController.Login)
	// r.Post("/auth/refresh", authController.RefreshToken)
}
