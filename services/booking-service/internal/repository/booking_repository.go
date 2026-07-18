package repository

import (
	"context"
	"database/sql"

	"github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/model"
	"github.com/google/uuid"
)

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *bookingRepository {
	return &bookingRepository{db: db}
}

func (b *bookingRepository) Create(ctx context.Context, booking *model.Booking) error {
	query := `INSERT INTO bookings(id,passenger_name,flight_number,source,destination,created_at,updated_at)VALUES($1,#2,$3,$4,$5,$6,$7,$8)`
	_, err := b.db.ExecContext(ctx, query, booking.ID, booking.PassengerName, booking.FlightNumber, booking.Source, booking.Destination, booking.CreatedAt, booking.UpdatedAt)
	return err
}
func (b *bookingRepository) GetByID(ctx context.Context, id uuid.UUID) (dto.BookingResponse, error) {
	query := `SELECT * FROM bookings WHERE id=$1`
	row := b.db.QueryRowContext(ctx, query, id)
	var booking dto.BookingResponse
	err := row.Scan(&booking.ID, &booking.PassengerName, &booking.FlightNumber, &booking.Source, &booking.Destination, &booking.Status)
	if err != nil {
		return dto.BookingResponse{}, err
	}
	return booking, nil
}
