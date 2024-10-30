// controllers/payments_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/models"
	"github.com/joshua468/ecommerce/services"
)

func ProcessPayment(c *gin.Context) {
	var paymentRequest models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payment input"})
		return
	}

	// Call ProcessPayment from services with models.PaymentRequest
	paymentResponse, err := services.ProcessPayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Payment process failed"})
		return
	}

	c.JSON(http.StatusOK, paymentResponse)
}
