package errors

import "net/http"

var (
	ErrInvalidRequest = &AppError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_REQUEST",
		Message:    "Invalid request",
	}

	ErrInternalServer = &AppError{
		HTTPStatus: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
	}

	ErrUnauthorized = &AppError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
		Message:    "Unauthorized",
	}
)
