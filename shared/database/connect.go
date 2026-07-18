package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RupeshMahanta1994/flightops/shared/config"
	_ "github.com/lib/pq"
)

func Init(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	//Create DB connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
		return nil, err
	}
	//Connection pool configuration
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)
	//Verify DB connection
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Failed to ping DB: %v", err)
		return nil, err
	}
	return db, nil
}
