package router

import (
	"training-plan-api/container"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(r fiber.Router, deps *container.AppDependencies) {

	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	// // Dashboard / stats
	// r.Get("/stats", deps.DepartmentController.GetStats)

	// Department management
	r.Post("/departments", deps.DepartmentController.Create)
	r.Put("/departments/:id", deps.DepartmentController.Update)
	r.Delete("/departments/:id", deps.DepartmentController.Delete)
	r.Get("/departments", deps.DepartmentController.FindPaginated)
	r.Get("/departments/:id", deps.DepartmentController.FindById)

	// Department List for assigning to users
	r.Get("/departments-list", deps.DepartmentController.GetDepartmentsList)

	// User management (Admin has full CRUD)
	r.Post("/users", deps.UserController.AdminCreate)
	r.Put("/users/:id", deps.UserController.AdminUpdate)
	r.Delete("/users/:id", deps.UserController.AdminDelete)
	r.Get("/users", deps.UserController.AdminFindAll)
	r.Get("/users/:id", deps.UserController.AdminFindById)

	// Course management
	r.Post("/courses", deps.CourseController.Create)
	r.Put("/courses/:id", deps.CourseController.Update)
	r.Delete("/courses/:id", deps.CourseController.Delete)
	r.Get("/courses", deps.CourseController.FindPaginated)
	r.Get("/courses/:id", deps.CourseController.FindById)

	// // Records
	// r.Get("/records", deps.RecordController.FindAllPaginated)

	// // Certificates (approval flow)
	r.Get("/certificates", deps.CertificateController.FindAllPending)
	r.Put("/certificates/:id/approve", deps.CertificateController.Approve)
	r.Put("/certificates/:id/reject", deps.CertificateController.Reject)
}
