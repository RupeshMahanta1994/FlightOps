package errors

import "net/http"

var (
	ErrPassengerNotFound = &AppError{
		HTTPStatus: http.StatusNotFound,
		Code:       "PASSENGER_NOT_FOUND",
		Message:    "Passenger not found",
	}

	ErrPassengerAlreadyCheckedIn = &AppError{
		HTTPStatus: http.StatusConflict,
		Code:       "PASSENGER_ALREADY_CHECKED_IN",
		Message:    "Passenger already checked in",
	}
)
