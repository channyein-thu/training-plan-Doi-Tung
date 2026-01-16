package router

import (
	"training-plan-api/controller"

	"github.com/gofiber/fiber/v2"
)

func DepartmentRoutes(api fiber.Router, controller *controller.DepartmentController) {
	departments := api.Group("/departments")

	departments.Post("/", controller.Create)
	departments.Get("/", controller.FindAll)
	departments.Get("/:departmentId", controller.FindById)
	departments.Put("/:departmentId", controller.Update)
	departments.Delete("/:departmentId", controller.Delete)
}

func CourseRoutes(api fiber.Router, courseController *controller.CourseController) {
	courses := api.Group("/courses")

	courses.Post("/", courseController.Create)
	courses.Put("/:courseId", courseController.Update)
	courses.Delete("/:courseId", courseController.Delete)
	courses.Get("/:courseId", courseController.FindById)
	courses.Get("/", courseController.FindPaginated)
}