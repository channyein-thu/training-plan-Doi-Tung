package controller

import (
	"strconv"
	"training-plan-api/data/request"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type CertificateController struct {
	certificateService service.CertificateService
}

func NewCertificateController(certificateService service.CertificateService) *CertificateController {
	return &CertificateController{
		certificateService: certificateService,
	}
}	


func (c *CertificateController) FindByCurrentUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	certificates, err := c.certificateService.FindByCurrentUser(userID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    certificates,
	})
}

func (c *CertificateController) Upload(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	// Parse multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		return helper.BadRequest("Invalid form data")
	}

	// Get training ID
	trainingIDStr := ctx.FormValue("trainingId")
	if trainingIDStr == "" {
		return helper.BadRequest("Training ID is required")
	}

	trainingID, err := strconv.ParseUint(trainingIDStr, 10, 32)
	if err != nil {
		return helper.BadRequest("Invalid training ID")
	}

	// Get file
	files := form.File["image"]
	if len(files) == 0 {
		return helper.BadRequest("Certificate image is required")
	}

	file := files[0]

	// Get optional description
	description := ctx.FormValue("description")
	var descPtr *string
	if description != "" {
		descPtr = &description
	}

	req := request.CreateCertificateRequest{
		TrainingID:   uint(trainingID),
		Image:       file.Filename,
		Description: descPtr,
	}

	if err := c.certificateService.Upload(userID, req, file); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Certificate uploaded successfully",
	})
}

func (c *CertificateController) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	certificateID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.BadRequest("Invalid certificate ID")
	}

	if err := c.certificateService.Delete(certificateID, userID); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Certificate deleted successfully",
	})
}
