package router

import (
	"training-plan-api/controller"
	"training-plan-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router, authController *controller.AuthController) {
	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	auth := r.Group("/auth")

	auth.Post("/admin/login", authController.AdminLogin)
	auth.Post("/manager/login", authController.ManagerLogin)
	auth.Post("/manager/register", authController.ManagerRegister)
	auth.Post("/staff/login", authController.StaffLogin)
	auth.Post("/staff/register", authController.StaffRegister)

	auth.Post("/refresh", authController.Refresh)

	auth.Post("/logout", authController.Logout)

	auth.Get("/me", middleware.JWTProtected, authController.GetMe)
}
