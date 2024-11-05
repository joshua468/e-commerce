package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/services"
)

type CreateOrderRequest struct {
	UserID uint `json:"user_id"` 
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cartItems, err := services.GetCartItems(req.UserID) // Use userID from request
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}

	// Check if cart is empty
	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cart is empty"})
		return
	}

	// Create order from cart items
	if err := services.CreateOrder(req.UserID, cartItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

func GetOrders(c *gin.Context) {
	// Parse userID from URL parameter
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve orders for the specified user
	orders, err := services.GetOrders(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	// If there are no orders, you may want to return a specific message or an empty list
	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No orders found for this user"})
		return
	}

	// Return the list of orders
	c.JSON(http.StatusOK, orders)
}
