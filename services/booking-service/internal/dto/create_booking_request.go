package dto

type CreateBookingRequest struct {
	PassengerName string `json:"passenger_name"`
	FlightNumber  string `json:"flight_number"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
}
