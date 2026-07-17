package main

import (
	"fmt"

	"github.com/RupeshMahanta1994/flightops/shared/config"
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
}
