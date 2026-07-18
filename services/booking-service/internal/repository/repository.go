package repository

import (
	"context"

	"github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/model"
	"github.com/google/uuid"
)

type BookingRepository interface {
	Create(ctx context.Context, bookings *model.Booking) error
	GetByID(ctx context.Context, id uuid.UUID) (dto.BookingResponse, error)
	Update(ctx context.Context, bookings *model.Booking) error
	Delete(ctx context.Context, id uuid.UUID) error
}
