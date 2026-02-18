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

	// Training Plan management
	r.Post("/training-plans", deps.TrainingPlanController.Create)
	r.Put("/training-plans/:trainingPlanId", deps.TrainingPlanController.Update)
	r.Delete("/training-plans/:trainingPlanId", deps.TrainingPlanController.Delete)
	r.Get("/training-plans", deps.TrainingPlanController.FindPaginated)
	r.Get("/training-plans/:trainingPlanId", deps.TrainingPlanController.FindById)

	// // Records
	// r.Get("/records", deps.RecordController.FindAllPaginated)

	// // Certificates (approval flow)
	r.Get("/certificates", deps.CertificateController.FindAllPending)
	r.Put("/certificates/:id/approve", deps.CertificateController.Approve)
	r.Put("/certificates/:id/reject", deps.CertificateController.Reject)
}
