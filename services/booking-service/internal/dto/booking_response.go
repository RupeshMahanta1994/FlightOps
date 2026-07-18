package dto

type BookingResponse struct {
	ID            string `json:"id"`
	PassengerName string `json:"passenger_name"`
	FlightNumber  string `json:"flight_number"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	Status        string `json:"status"`
}
