package router

import (
	"training-plan-api/container"

	"github.com/gofiber/fiber/v2"
)

func ManagerRoutes(r fiber.Router, deps *container.AppDependencies) {

	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	// User management (Manager can only create users in their department)
	r.Post("/users", deps.UserController.ManagerCreate)
	r.Get("/users", deps.UserController.ManagerFindDepartmentUsers)

	// // Department (with staff list)
	// r.Get("/departments/:id", deps.DepartmentController.FindByIdWithStaff)

	// // Register staff to course (YOUR FORM)
	// r.Post(
	// 	"/courses/:courseId/registrations",
	// 	deps.RecordController.RegisterStaff,
	// )

	// // Records
	// r.Get("/records", deps.RecordController.FindByDepartment)
	// r.Get("/records/:id", deps.RecordController.FindById)
	// r.Put("/records/:id", deps.RecordController.Update)
	// r.Delete("/records/:id", deps.RecordController.Delete)
}
