package router

import (
	"training-plan-api/container"
	"training-plan-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, deps *container.AppDependencies) {
	api := app.Group("/api/v1")

	// Public routes
	

	app.Get("/auth/google/login", deps.AuthOAuthController.GoogleLogin)
	app.Post("/auth/google/exchange", deps.AuthOAuthController.GoogleExchange)
	app.Post("/user/complete-profile", middleware.JWTProtected, deps.UserController.CompleteProfile)
    api.Get("/departments-list", deps.DepartmentController.GetDepartmentsList)
	
	AuthRoutes(api, deps.AuthController, deps.AuthOAuthController)
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
		api.Group("/staff", middleware.JWTProtected, middleware.RequireProfileComplete(deps.UserRepository) ),
		deps,
	)
}
