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
	if err != nil {
		log.Error("Error connecting to DB", err)
		db.Close()
	}
	log.Info("Databse connection successful", db.Stats().OpenConnections)
	defer db.Close()

	http.HandleFunc("/success", successHandler(log))

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
