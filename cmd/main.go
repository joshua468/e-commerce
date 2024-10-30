package main

import (
	"github.com/joho/godotenv"
	"github.com/joshua468/ecommerce/config"    // Ensure correct import paths
	migrations"github.com/joshua468/ecommerce/migration" // Ensure correct import paths
	"github.com/joshua468/ecommerce/routes"     // Ensure correct import paths
	"log"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize configurations
	config.ConnectDatabase()
	config.ConnectRedis()
	config.InitFlutterwave()

	// Run migrations
	migrations.Migrate()

	// Setup and run the server
	r := routes.SetupRouter()
	r.Run(":8080") // Start the server on port 8080
}
