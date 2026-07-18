package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID            uuid.UUID
	PassengerName string
	FlightNumber  string
	Source        string
	Destination   string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
