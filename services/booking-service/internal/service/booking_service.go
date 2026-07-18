package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/model"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/repository"
	"github.com/google/uuid"
)

type bookingService struct {
	repo repository.BookingRepository
	log  *slog.Logger
}

func NewBookingService(repo repository.BookingRepository, log *slog.Logger) *bookingService {
	return &bookingService{
		repo: repo,
		log:  log,
	}
}

func (s *bookingService) CreateBooking(passengerName, flightNumber, source, destination string) (*dto.BookingResponse, error) {
	booking := &model.Booking{
		ID:            uuid.New(),
		PassengerName: passengerName,
		FlightNumber:  flightNumber,
		Source:        source,
		Destination:   destination,
		Status:        "CONFIRMED",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	s.log.Info("persisting booking", "booking_id", booking.ID.String(), "flight_number", booking.FlightNumber)
	err := s.repo.Create(context.Background(), booking)
	if err != nil {
		s.log.Error("repository create failed", "booking_id", booking.ID.String(), "error", err.Error())
		return nil, err
	}
	return &dto.BookingResponse{
		ID:            booking.ID.String(),
		PassengerName: booking.PassengerName,
		FlightNumber:  booking.FlightNumber,
		Source:        booking.Source,
		Destination:   booking.Destination,
		Status:        booking.Status,
	}, nil
}

func (s *bookingService) GetBookingByID(id string) (dto.BookingResponse, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.log.Warn("invalid booking id", "booking_id", id)
		return dto.BookingResponse{}, err
	}

	s.log.Info("retrieving booking", "booking_id", id)
	return s.repo.GetByID(context.Background(), uuid)
}
