package controller

import (
	"strconv"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type RecordController struct {
	service service.RecordService
}

func NewRecordController(service service.RecordService) *RecordController {
	return &RecordController{service: service}
}

func (c *RecordController) RegisterStaff(ctx *fiber.Ctx) error {
	courseId, err := strconv.Atoi(ctx.Params("courseId"))
	if err != nil {
		return helper.BadRequest("Invalid course ID")
	}

	var req request.RegisterStaffRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	if err := c.service.RegisterStaff(uint(courseId), req); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Staff registered to course successfully",
	})
}
func (c *RecordController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid record ID")
	}

	result, err := c.service.FindById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Status: "SUCCESS",
		Data:   result,
	})
}

func (c *RecordController) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid record ID")
	}

	var req request.UpdateRecordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	if err := c.service.Update(id, req); err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Record updated successfully",
	})
}

func (c *RecordController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid record ID")
	}

	if err := c.service.Delete(id); err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Record deleted successfully",
	})
}
