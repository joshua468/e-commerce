package services

import (
	"github.com/joshua468/ecommerce/config"
	"github.com/joshua468/ecommerce/models"
)

// CreateOrder creates a new order based on the user's cart items
func CreateOrder(userID uint, cartItems []models.Cart) error {
	order := models.Order{
		UserID: userID,
		Status: "Pending",
	}

	// Begin a transaction
	tx := config.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Add each cart item to the order
	for _, item := range cartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetOrders retrieves all orders for a specified user
func GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderItems retrieves items for a given order ID
func GetOrderItems(orderID uint) ([]models.OrderItem, error) {
	var items []models.OrderItem
	if err := config.DB.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
