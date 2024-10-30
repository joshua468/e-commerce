package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/models"
	"github.com/joshua468/ecommerce/services"
	
)

func RegisterUser(c *gin.Context) {
var user models.User
if err:= c.ShouldBindJSON(&user);err!= nil{ 
c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid input"})
return
}
if err:= services.RegisterUser(&user);err!= nil {
	c.JSON(http.StatusInternalServerError,gin.H{"message":"User registered successfully"})
}
c.JSON(http.StatusCreated,gin.H{"message":"User registered"})
}

func LoginUser(c *gin.Context) {
var user models.User
if err:= c.ShouldBindJSON(&user);err!=nil {
	c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid input"})
	return
}
token,err:= services.LoginUser(user.Username,user.Password);
if err!=nil {
c.JSON(http.StatusUnauthorized,gin.H{"message":err.Error()})
return
}
c.JSON(http.StatusOK,gin.H{"messsage":"Login successfully","token":token})
return
}

func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"message":"Logout successfully"})
}