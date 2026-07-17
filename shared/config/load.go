package config

// 1. Read .env
// 2. Validate requred variables
// 3. Return Config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {
	//LOad .env file
	//Ignore error bcoz ind prod the variables will come from else where
	godotenv.Load()
	cfg := Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "flightops"),
			Port: getEnv("PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getRequiredEnv("DB_HOST"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getRequiredEnv("DB_USER"),
			Password: getRequiredEnv("DB_PASSWORD"),
			Name:     getRequiredEnv("DB_NAME"),
		},
		RabbitMQ: RabbitMQConfig{
			URL: getRequiredEnv("RABBITMQ_URL"),
		},
	}
	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variables %s is not set", key)
	}
	return value
}
