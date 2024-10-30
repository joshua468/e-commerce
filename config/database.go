package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global variables for configuration
var (
	DB             *gorm.DB
	FlutterwaveKey string
	FlutterwaveURL = "https://api.flutterwave.com/v3/charges?type=payment"
)

// ConnectDatabase initializes the database connection.
func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
		return
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connection successful")
}

// InitFlutterwave initializes Flutterwave configurations.
func InitFlutterwave() {
	FlutterwaveKey = os.Getenv("FLUTTERWAVE_KEY")
	if FlutterwaveKey == "" {
		log.Fatal("FLUTTERWAVE_KEY environment variable is not set")
	} else {
		log.Println("Flutterwave configured successfully")
	}
}
