package helper

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AppError struct {
	StatusCode int
	Message    string
	Errors     map[string]string `json:"errors,omitempty"`
}


func (e *AppError) Error() string {
	return e.Message
}

func ValidationError(errors map[string]string) error {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "Validation failed",
		Errors:     errors,
	}
}
func FormatValidationError(err error) map[string]string {
	out := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return out
	}

	for _, e := range validationErrors {
		field := strings.ToLower(e.Field())

		switch e.Tag() {
		case "required":
			out[field] = "is required"
		case "min":
			out[field] = "must be at least " + e.Param() + " characters"
		case "gte":
			out[field] = "must be greater than or equal to " + e.Param()
		default:
			out[field] = "is invalid"
		}
	}

	return out
}

// Helpers
func BadRequest(msg string) error {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NotFound(msg string) error {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func Internal(msg string) error {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

func Unauthorized(msg string) error {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
	}
}

func Forbidden(msg string) error {
	return &AppError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
	}
}

func InternalServerError(msg string) error {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}
