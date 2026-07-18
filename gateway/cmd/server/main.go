package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/RupeshMahanta1994/flightops/shared/config"
	"github.com/RupeshMahanta1994/flightops/shared/database"
	"github.com/RupeshMahanta1994/flightops/shared/errors"
	"github.com/RupeshMahanta1994/flightops/shared/logger"
	"github.com/RupeshMahanta1994/flightops/shared/response"
)

func main() {
	cfg := config.Load()
	fmt.Println("Application Configuration")
	fmt.Println("--------------------------")

	fmt.Println("App Name       :", cfg.App.Name)
	fmt.Println("App Port       :", cfg.App.Port)

	fmt.Println("DB Host        :", cfg.Database.Host)
	fmt.Println("DB Port        :", cfg.Database.Port)
	fmt.Println("DB User        :", cfg.Database.User)
	fmt.Println("DB Name        :", cfg.Database.Name)

	fmt.Println("RabbitMQ URL   :", cfg.RabbitMQ.URL)

	log := logger.New("gateway")
	log.Info("Application Server started")
	db, err := database.Init(cfg)
	fmt.Println("Databse connection successful", db.Stats().OpenConnections, err)

	http.HandleFunc("/success", successHandler(log))
	// http.HandleFunc("/bad-request", badRequestHandler(log))
	// http.HandleFunc("/not-found", notFoundHandler(log))
	// http.HandleFunc("/internal-error", internalErrorHandler(log))

	log.Info("Gateway started", slog.String("port", cfg.App.Port))

	if err := http.ListenAndServe(":"+cfg.App.Port, nil); err != nil {
		log.Error("failed to start server", "error", err)
	}
}

func successHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Error(w, errors.ErrBookingNotFound)
		response.Success(w, 200, "Success response", 102)
		log.Info("Successfully sent the response")
	}
}

// func badRequestHandler(log *slog.Logger) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := response.BadRequest(
// 			w,
// 			"204",
// 			"Passenger ID is required",
// 		); err != nil {
// 			log.Error("failed to write response", "error", err)
// 		}
// 	}
// }

// func notFoundHandler(log *slog.Logger) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := response.NotFound(
// 			w,
// 			"BOOKING_NOT_FOUND",
// 			"Booking does not exist",
// 		); err != nil {
// 			log.Error("failed to write response", "error", err)
// 		}
// 	}
// }

// func internalErrorHandler(log *slog.Logger) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := response.InternalServerError(
// 			w,
// 			"DATABASE_ERROR",
// 			"Unable to connect to database",
// 		); err != nil {
// 			log.Error("failed to write response", "error", err)
// 		}
// 	}

// }
