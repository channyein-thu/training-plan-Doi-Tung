package helper

import "net/http"

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
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
