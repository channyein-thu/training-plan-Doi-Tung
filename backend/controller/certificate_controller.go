package controller

import (
	"strconv"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type CertificateController struct {
	service service.CertificateService
}

func NewCertificateController(service service.CertificateService) *CertificateController {
	return &CertificateController{service: service}
}

// ================= USER =================

func (c *CertificateController) FindByCurrentUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	result, err := c.service.FindByCurrentUser(userID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   result,
	})
}

func (c *CertificateController) Upload(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	trainingID, err := strconv.Atoi(ctx.FormValue("trainingId"))
	if err != nil {
		return helper.BadRequest("Invalid training ID")
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return helper.BadRequest("Certificate image is required")
	}

	description := ctx.FormValue("description")
	var desc *string
	if description != "" {
		desc = &description
	}

	req := request.CreateCertificateRequest{
		TrainingID:   uint(trainingID),
		Description: desc,
	}

	if err := c.service.Upload(userID, req, file); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Certificate uploaded successfully",
	})
}

func (c *CertificateController) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid certificate ID")
	}

	if err := c.service.Delete(id, userID); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Certificate deleted successfully",
	})
}

// ================= ADMIN =================

func (c *CertificateController) FindAllPending(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.service.FindAllPending(page, limit)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   result,
	})
}

func (c *CertificateController) Approve(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid certificate ID")
	}

	if err := c.service.Approve(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Certificate approved successfully",
	})
}

func (c *CertificateController) Reject(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid certificate ID")
	}

	if err := c.service.Reject(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Certificate rejected successfully",
	})
}
