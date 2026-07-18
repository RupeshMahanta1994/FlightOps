package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/RupeshMahanta1994/flightops/booking-service/internal/dto"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/model"
	"github.com/google/uuid"
)

type bookingRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewBookingRepository(db *sql.DB, log *slog.Logger) *bookingRepository {
	return &bookingRepository{db: db, log: log}
}

func (b *bookingRepository) Create(ctx context.Context, booking *model.Booking) error {
	query := `INSERT INTO bookings(id, passenger_name, flight_number, source, destination, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	b.log.Info("Creating booking")
	_, err := b.db.ExecContext(ctx, query, booking.ID, booking.PassengerName, booking.FlightNumber, booking.Source, booking.Destination, booking.Status, booking.CreatedAt, booking.UpdatedAt)
	if err != nil {
		b.log.Error("Failed to create booking", "error", err.Error())
		return err
	}
	b.log.Info("Booking created successfully")
	return err
}
func (b *bookingRepository) GetByID(ctx context.Context, id uuid.UUID) (dto.BookingResponse, error) {
	b.log.Info("Fetching booking by ID", "id", id)
	query := `SELECT * FROM bookings WHERE id=$1`
	row := b.db.QueryRowContext(ctx, query, id)
	var booking model.Booking
	err := row.Scan(&booking.ID, &booking.PassengerName, &booking.FlightNumber, &booking.Source, &booking.Destination, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
	if err != nil {
		b.log.Error("Failed to fetch booking", "error", err.Error())
		return dto.BookingResponse{}, err
	}
	var response dto.BookingResponse
	response.ID = booking.ID.String()
	response.PassengerName = booking.PassengerName
	response.FlightNumber = booking.FlightNumber
	response.Source = booking.Source
	response.Destination = booking.Destination
	response.Status = booking.Status
	if err != nil {
		b.log.Error("Failed to scan booking", "error", err.Error())
		return dto.BookingResponse{}, err
	}
	return response, nil

}
func (b *bookingRepository) Update(ctx context.Context, booking *model.Booking) error {
	query := `UPDATE bookings SET passenger_name=$1, flight_number=$2, source=$3, destination=$4, status=$5, updated_at=$6 WHERE id=$7`
	b.log.Info("Updating booking", "id", booking.ID)
	_, err := b.db.ExecContext(ctx, query, booking.PassengerName, booking.FlightNumber, booking.Source, booking.Destination, booking.Status, booking.UpdatedAt, booking.ID)
	if err != nil {
		b.log.Error("Failed to update booking", "error", err.Error())
		return err
	}
	b.log.Info("Booking updated successfully")
	return nil
}

func (b *bookingRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM bookings WHERE id=$1`
	b.log.Info("Deleting booking", "id", id)
	_, err := b.db.ExecContext(ctx, query, id)
	if err != nil {
		b.log.Error("Failed to delete booking", "error", err.Error())
		return err
	}
	b.log.Info("Booking deleted successfully")
	return nil
}

func (b *bookingRepository) ListAllBookings(ctx context.Context) ([]dto.BookingResponse, error) {
	query := `SELECT * FROM bookings`
	b.log.Info("Listing all bookings")
	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		b.log.Error("Failed to list bookings", "error", err.Error())
		return nil, err
	}
	defer rows.Close()

	var bookings []dto.BookingResponse
	for rows.Next() {
		var booking model.Booking
		err := rows.Scan(&booking.ID, &booking.PassengerName, &booking.FlightNumber, &booking.Source, &booking.Destination, &booking.Status, &booking.CreatedAt, &booking.UpdatedAt)
		if err != nil {
			b.log.Error("Failed to scan booking", "error", err.Error())
			return nil, err
		}
		bookings = append(bookings, dto.BookingResponse{
			ID:            booking.ID.String(),
			PassengerName: booking.PassengerName,
			FlightNumber:  booking.FlightNumber,
			Source:        booking.Source,
			Destination:   booking.Destination,
			Status:        booking.Status,
		})
	}
	if err := rows.Err(); err != nil {
		b.log.Error("Error iterating over bookings", "error", err.Error())
		return nil, err
	}
	b.log.Info("Bookings listed successfully", "count", len(bookings))
	return bookings, nil
}
