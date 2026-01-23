package controller

import (
	"strconv"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type CourseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) *CourseController {
	return &CourseController{courseService: courseService}
}

// CREATE
func (c *CourseController) Create(ctx *fiber.Ctx) error {
	var req request.CreateCourseRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid course data")
	}

	if err := c.courseService.Create(req); err != nil {
		return err // handled by global error handler
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Course created successfully",
	})
}

// UPDATE
func (c *CourseController) Update(ctx *fiber.Ctx) error {
	var req request.UpdateCourseRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid course data")
	}

	id, err := strconv.Atoi(ctx.Params("courseId"))
	if err != nil {
		return helper.BadRequest("Invalid course ID")
	}

	if err := c.courseService.Update(id, req); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Course updated successfully",
	})
}

// DELETE
func (c *CourseController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("courseId"))
	if err != nil {
		return helper.BadRequest("Invalid course ID")
	}

	if err := c.courseService.Delete(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Course deleted successfully",
	})
}

// FIND BY ID
func (c *CourseController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("courseId"))
	if err != nil {
		return helper.BadRequest("Invalid course ID")
	}

	course, err := c.courseService.FindById(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Course retrieved successfully",
		Data:    course,
	})
}

// FIND PAGINATED
func (c *CourseController) FindPaginated(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.courseService.FindPaginated(page, limit)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Courses retrieved successfully",
		Data:    result,
	})
}
