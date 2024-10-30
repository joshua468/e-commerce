// controllers/order_controller.go
package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/joshua468/ecommerce/services"
)

func CreateOrder(c *gin.Context) {
    // Parse userID from URL parameter
    userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    cartItems, err := services.GetCartItems(uint(userID)) // Ensure you have a function to get cart items
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
    if err := services.CreateOrder(uint(userID), cartItems); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
// GetOrders fetches all orders for a specific user
func GetOrders(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    orders, err := services.GetOrders(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}
