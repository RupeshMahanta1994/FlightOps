package service

import "github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"

type BookingService interface {
	CreateBooking(passengerName, flightNumber, source, destination string) (*dto.BookingResponse, error)
	GetBookingByID(id string) (dto.BookingResponse, error)
	UpdateBooking(id, passengerName, flightNumber, source, destination, status string) (*dto.BookingResponse, error)
	DeleteBooking(id string) error
}
