package controller

import (
	"fmt"
	"strconv"
	"time"
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
	trainingPlanId, err := strconv.Atoi(ctx.Params("trainingPlanId"))
	if err != nil {
		return helper.BadRequest("Invalid training plan ID")
	}

	var req request.RegisterStaffRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	if err := c.service.RegisterStaff(uint(trainingPlanId), req); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Staff registered to training plan successfully",
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

func (c *RecordController) FindRecordByCurrentDepartment(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.service.FindByManager(userID, page, limit)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Records retrieved successfully",
		Data:    result,
	})

}

func (c *RecordController) FindByCurrentUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.service.FindByUser(userID, page, limit)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Records retrieved successfully",
		Data:    result,
	})
}

func (c *RecordController) Search(ctx *fiber.Ctx) error {

	var req request.RecordFilterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	result, err := c.service.Search(req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Records fetched successfully",
		Data:    result,
	})
}

func (c *RecordController) Export(ctx *fiber.Ctx) error {

	var req request.RecordFilterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	file, err := c.service.Export(req)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("records_%d.xlsx", time.Now().Unix())

	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", "attachment; filename="+fileName)

	return file.Write(ctx.Response().BodyWriter())
}
