// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/ecommerce/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.POST("/logout", controllers.LogoutUser)
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetAvailableProducts)
		api.POST("/orders", controllers.CreateOrder)
		api.GET("/orders/:user_id", controllers.GetOrders) // Ensure this is included
		api.POST("/cart", controllers.AddToCart)
		api.GET("/cart/:user_id", controllers.GetCartItems)
		api.DELETE("/cart/:cart_id", controllers.RemoveFromCart)
		api.POST("/payment", controllers.ProcessPayment)
	}
	return r
}
