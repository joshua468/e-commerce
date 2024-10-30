package migrations

import (
	"log"

	"github.com/joshua468/ecommerce/config"
	"github.com/joshua468/ecommerce/models"
)

func Migrate() {
	// Use AutoMigrate for GORM-managed tables
	if err := config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Payment{}, &models.Shop{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	createTables := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS shops (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		name VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		shop_id INT REFERENCES shops(id) ON DELETE CASCADE,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		price DECIMAL(10, 2) NOT NULL,
		stock INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS carts (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		product_id INT REFERENCES products(id) ON DELETE CASCADE,
		quantity INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		status VARCHAR(20) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS order_items (
		id SERIAL PRIMARY KEY,
		order_id INT REFERENCES orders(id) ON DELETE CASCADE,
		product_id INT REFERENCES products(id) ON DELETE CASCADE,
		quantity INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS payments (
		id SERIAL PRIMARY KEY,
		order_id INT REFERENCES orders(id) ON DELETE CASCADE,
		amount DECIMAL(10, 2) NOT NULL,
		payment_status VARCHAR(20) NOT NULL,
		transaction_id VARCHAR(100) UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	result := config.DB.Exec(createTables)
	if result.Error != nil {
		log.Fatalf("Failed to create tables: %v", result.Error)
	}

	log.Println("Database tables created successfully")
}
