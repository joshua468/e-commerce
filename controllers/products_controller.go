package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/models"
	"github.com/joshua468/ecommerce/services"
)

func CreateProduct(c *gin.Context)  {
var products models.Product
if err:=c.ShouldBindJSON(&products);err!=nil {
	c.JSON(http.StatusBadRequest,gin.H{"message":"invalid product input"})
	return
}
if err:= services.CreateProduct(&products);err!=nil  {
	c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
	return
}
c.JSON(http.StatusCreated,gin.H{"message":"products  created successfully"})
}

func GetAvailableProducts(c *gin.Context) {
	products,err := services.GetAvailableProducts()
	if err!= nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}
	c.JSON(http.StatusOK,products)
}