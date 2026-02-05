package router

import (
	"training-plan-api/container"
	"training-plan-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, deps *container.AppDependencies) {
	api := app.Group("/api/v1")

	// Public routes
	AuthRoutes(api, deps.AuthController)

	// Role-based routes with JWT and role middleware
	AdminRoutes(
		api.Group("/admin", middleware.JWTProtected, middleware.AdminOnly),
		deps,
	)

	ManagerRoutes(
		api.Group("/manager", middleware.JWTProtected, middleware.ManagerOnly),
		deps,
	)

	StaffRoutes(
		api.Group("/staff", middleware.JWTProtected, middleware.StaffOnly),
		deps,
	)
}
