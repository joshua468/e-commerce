// models/order.go
package models

import (
	"time"
)

type Order struct {
	ID        uint        `json:"id" gorm:"primary_key"`
	UserID    uint        `json:"user_id"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint `json:"id" gorm:"primary_key"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
