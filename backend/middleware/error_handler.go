package middleware

import (
	"errors"
	"net/http"
	"training-plan-api/data/response"
	"training-plan-api/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Default error
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Validation errors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: ve.Error(),
		})
	}

	// AppError (custom)
	var appErr *helper.AppError
	if errors.As(err, &appErr) {
		return ctx.Status(appErr.StatusCode).JSON(response.Response{
			Code:    appErr.StatusCode,
			Status:  http.StatusText(appErr.StatusCode),
			Message: appErr.Message,
		})
	}

	// Fiber error (404, etc.)
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return ctx.Status(fiberErr.Code).JSON(response.Response{
			Code:    fiberErr.Code,
			Status:  http.StatusText(fiberErr.Code),
			Message: fiberErr.Message,
		})
	}

	// Fallback
	return ctx.Status(code).JSON(response.Response{
		Code:    code,
		Status:  "ERROR",
		Message: message,
	})
}
