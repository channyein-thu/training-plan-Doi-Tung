package middleware

import (
	"net/http"
	"training-plan-api/data/response"
	"training-plan-api/helper"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	// Handle custom AppError
	if appErr, ok := err.(*helper.AppError); ok {
		return ctx.Status(appErr.StatusCode).JSON(response.Response{
			Status:  http.StatusText(appErr.StatusCode),
			Message: appErr.Message,
			Data:    appErr.Errors,
		})
	}

	// Handle Fiber errors (404, 405, etc.)
	if fiberErr, ok := err.(*fiber.Error); ok {
		return ctx.Status(fiberErr.Code).JSON(response.Response{
			Status:  http.StatusText(fiberErr.Code),
			Message: fiberErr.Message,
		})
	}

	// Handle unexpected errors
	return ctx.Status(fiber.StatusInternalServerError).JSON(response.Response{
		Status:  "ERROR",
		Message: "Internal server error",
	})
}
