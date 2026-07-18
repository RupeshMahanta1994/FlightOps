package main

// This file contains the main function for the booking service server.

import (
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/handler"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/repository"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/router"
	"github.com/RupeshMahanta1994/flightops/booking-service/internal/service"
	"github.com/RupeshMahanta1994/flightops/shared/config"
	"github.com/RupeshMahanta1994/flightops/shared/database"
	"github.com/RupeshMahanta1994/flightops/shared/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	log := logger.New("Booking-Service")
	db, err := database.Init(cfg)
	if err != nil {
		panic(err)
	}
	log.Info("Database connection successful")
	// Initialize the booking repository
	repo := repository.NewBookingRepository(db)

	// Initialize the booking service
	bookingService := service.NewBookingService(repo, log)

	// Initialize the booking handler
	bookingHandler := handler.NewBookingHandler(bookingService, log)

	//rotes are defined in rotuer folder
	app := fiber.New()
	bookingRouter := router.NewBookingRouter(app, bookingHandler)
	bookingRouter.SetupRoutes()

	log.Info("Starting server on port 8082")
	if err := app.Listen(":8082"); err != nil {
		log.Error("Failed to start server: ", err)
	}
}
