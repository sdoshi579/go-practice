package errs

import "net/http"


type AppError struct {
	Code int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}