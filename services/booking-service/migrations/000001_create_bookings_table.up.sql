CREATE TABLE bookings (
    id UUID PRIMARY KEY,
    passenger_name VARCHAR(255) NOT NULL,
    flight_number VARCHAR(20) NOT NULL,
    source VARCHAR(10) NOT NULL,
    destination VARCHAR(10) NOT NULL,
    status VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);