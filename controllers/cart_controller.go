package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/models"
	"github.com/joshua468/ecommerce/services"
)

func AddToCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddToCart(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetCartItems(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
	cartItems, err := services.GetCartItems(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

func RemoveFromCart(c *gin.Context) {
	cartID, _ := strconv.ParseUint(c.Param("cart_id"), 10, 32)
	if err := services.RemoveFromCart(uint(cartID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
