package config

// Example responsibilities:

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	RabbitMQ RabbitMQConfig
}

// AppConfig struct

type AppConfig struct {
	Name string
	Port string
}

// DatabaseConfig struct
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// RabbitMQConfig struct
type RabbitMQConfig struct {
	URL string
}
