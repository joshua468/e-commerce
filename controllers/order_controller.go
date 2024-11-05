package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/services"
)

type CreateOrderRequest struct {
	UserID uint `json:"user_id"` 
}

type OrdersResponse struct {
	Message string      `json:"message"`
	Orders  []OrderDetail `json:"orders"`
}

type OrderDetail struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"` 
	UpdatedAt string `json:"updated_at"` 
	Items     []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        uint `json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"` // Change to int
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cartItems, err := services.GetCartItems(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cart is empty"})
		return
	}

	if err := services.CreateOrder(req.UserID, cartItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

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

	if len(orders) == 0 {
		c.JSON(http.StatusOK, OrdersResponse{
			Message: "No orders found for this user",
			Orders:  []OrderDetail{},
		})
		return
	}

	var responseOrders []OrderDetail
	for _, order := range orders {
		orderItems, err := services.GetOrderItems(order.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order items"})
			return
		}

		var mappedItems []OrderItem
		for _, item := range orderItems {
			mappedItems = append(mappedItems, OrderItem{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity, // This is now safe
			})
		}

		responseOrders = append(responseOrders, OrderDetail{
			ID:        order.ID,
			UserID:    order.UserID,
			Status:    order.Status,
			CreatedAt: order.CreatedAt.Format("2006-01-02T15:04:05Z"), 
			UpdatedAt: order.UpdatedAt.Format("2006-01-02T15:04:05Z"), 
			Items:     mappedItems,
		})
	}

	c.JSON(http.StatusOK, OrdersResponse{
		Message: "Orders retrieved successfully",
		Orders:  responseOrders,
	})
}
