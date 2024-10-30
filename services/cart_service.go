package services

import (
	"github.com/joshua468/ecommerce/config"
	"github.com/joshua468/ecommerce/models"
)

func AddToCart(cart *models.Cart) error {
	if err := config.DB.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

func GetCartItems(userID uint) ([]models.Cart, error) {
	var cartItems []models.Cart
	if err := config.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func RemoveFromCart(cartID uint) error {
	if err := config.DB.Delete(&models.Cart{}, cartID).Error; err != nil {
		return err
	}
	return nil
}
