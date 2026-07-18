package errors

import "net/http"

var (
	ErrBookingNotFound = &AppError{
		HTTPStatus: http.StatusNotFound,
		Code:       "BOOKING_NOT_FOUND",
		Message:    "Booking not found",
	}

	ErrBookingAlreadyExists = &AppError{
		HTTPStatus: http.StatusConflict,
		Code:       "BOOKING_ALREADY_EXISTS",
		Message:    "Booking already exists",
	}

	ErrFlightClosed = &AppError{
		HTTPStatus: http.StatusConflict,
		Code:       "FLIGHT_CLOSED",
		Message:    "Flight is already closed",
	}
)
