package controller

import (
	"strconv"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type DepartmentController struct {
	departmentService service.DepartmentService
}

func NewDepartmentController(departmentService service.DepartmentService) *DepartmentController {
	return &DepartmentController{
		departmentService: departmentService,
	}
}

func (c *DepartmentController) Create(ctx *fiber.Ctx) error {
	var req request.CreateDepartmentRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid department data")
	}

	if err := c.departmentService.Create(req); err != nil {
		return err // 🔥 handled by global error handler
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Code:    fiber.StatusCreated,
		Status:  "SUCCESS",
		Message: "Department created successfully",
	})
}

func (c *DepartmentController) Update(ctx *fiber.Ctx) error {
	var req request.UpdateDepartmentRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid department data")
	}

	id, err := strconv.Atoi(ctx.Params("departmentId"))
	if err != nil {
		return helper.BadRequest("Invalid department ID")
	}

	if err := c.departmentService.Update(id, req); err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Department updated successfully",
	})
}

func (c *DepartmentController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("departmentId"))
	if err != nil {
		return helper.BadRequest("Invalid department ID")
	}

	if err := c.departmentService.Delete(id); err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Department deleted successfully",
	})
}

func (c *DepartmentController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("departmentId"))
	if err != nil {
		return helper.BadRequest("Invalid department ID")
	}

	department, err := c.departmentService.FindById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Department retrieved successfully",
		Data:    department,
	})
}

func (c *DepartmentController) FindAll(ctx *fiber.Ctx) error {
	departments, err := c.departmentService.FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Code:    fiber.StatusOK,
		Status:  "SUCCESS",
		Message: "Departments retrieved successfully",
		Data:    departments,
	})
}
