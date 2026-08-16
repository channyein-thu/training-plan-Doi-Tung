package controller

import (
	"strconv"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type TrainingPlanController struct {
	trainingPlanService service.TrainingPlanService
}

func NewTrainingPlanController(trainingPlanService service.TrainingPlanService) *TrainingPlanController {
	return &TrainingPlanController{trainingPlanService: trainingPlanService}
}

// CREATE
func (c *TrainingPlanController) Create(ctx *fiber.Ctx) error {
	var req request.CreateTrainingPlanRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid training plan data")
	}

	if err := c.trainingPlanService.Create(req); err != nil {
		return err // handled by global error handler
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Training plan created successfully",
	})
}

// UPDATE
func (c *TrainingPlanController) Update(ctx *fiber.Ctx) error {
	var req request.UpdateTrainingPlanRequest

	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid training plan data")
	}

	id, err := strconv.Atoi(ctx.Params("trainingPlanId"))
	if err != nil {
		return helper.BadRequest("Invalid training plan ID")
	}

	if err := c.trainingPlanService.Update(id, req); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Training plan updated successfully",
	})
}

// DELETE
func (c *TrainingPlanController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("trainingPlanId"))
	if err != nil {
		return helper.BadRequest("Invalid training plan ID")
	}

	if err := c.trainingPlanService.Delete(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Training plan deleted successfully",
	})
}

// FIND BY ID
func (c *TrainingPlanController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("trainingPlanId"))
	if err != nil {
		return helper.BadRequest("Invalid training plan ID")
	}

	trainingPlan, err := c.trainingPlanService.FindById(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Training plan retrieved successfully",
		Data:    trainingPlan,
	})
}

// FIND PAGINATED
func (c *TrainingPlanController) FindPaginated(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.trainingPlanService.FindPaginated(page, limit)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Training plans retrieved successfully",
		Data:    result,
	})
}
