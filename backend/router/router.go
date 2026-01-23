package router

import (
	"training-plan-api/container"

	"github.com/gofiber/fiber/v2"
)


func RegisterRoutes(app *fiber.App, deps *container.AppDependencies) {
	api := app.Group("/api/v1")

	// Public routes
	AuthRoutes(api, deps.AuthController)

	// Protected routes
	// api.Use(middleware.JWTProtected)

	// Role-based routes
	AdminRoutes(
		api.Group("/admin", ),
		deps,
	)

	ManagerRoutes(
		api.Group("/manager", ),
		// api.Group("/manager", middleware.RequireRole("MANAGER")),
		deps,
	)

	StaffRoutes(
		api.Group("/staff", ),
		deps,
	)
}

// func DepartmentRoutes(api fiber.Router, controller *controller.DepartmentController) {
// 	departments := api.Group("/departments")

// 	departments.Post("/", controller.Create)
// 	departments.Get("/", controller.FindPaginated)
// 	departments.Get("/:departmentId", controller.FindById)
// 	departments.Put("/:departmentId", controller.Update)
// 	departments.Delete("/:departmentId", controller.Delete)
// }

// func CourseRoutes(api fiber.Router, courseController *controller.CourseController) {
// 	courses := api.Group("/courses")

// 	courses.Post("/", courseController.Create)
// 	courses.Put("/:courseId", courseController.Update)
// 	courses.Delete("/:courseId", courseController.Delete)
// 	courses.Get("/:courseId", courseController.FindById)
// 	courses.Get("/", courseController.FindPaginated)
// }

// func AuthRoutes(api fiber.Router, authController *controller.AuthController) {
// 	// Admin routes
// 	admin := api.Group("/admin")
// 	admin.Post("/login", authController.AdminLogin)

// 	// Manager routes
// 	manager := api.Group("/manager")
// 	manager.Post("/login", authController.ManagerLogin)
// 	manager.Post("/register", authController.ManagerRegister)

// 	// Staff routes
// 	staff := api.Group("/staff")
// 	staff.Post("/login", authController.StaffLogin)
// 	staff.Post("/register", authController.StaffRegister)

// 	// Protected route - get current user
// 	api.Get("/me", middleware.JWTProtected, authController.GetMe)
// }
