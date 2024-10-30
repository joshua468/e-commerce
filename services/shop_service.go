package services

import (
	"github.com/joshua468/ecommerce/models"
	"github.com/joshua468/ecommerce/config"
)

func CreateShop(shop *models.Shop) error {
	if err := config.DB.Create(shop).Error; err != nil {
		return err
	}
	return nil
}

func GetShopByID(shopID uint) (models.Shop, error) {
	var shop models.Shop
	if err := config.DB.Preload("Products").First(&shop, shopID).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func GetAllShops() ([]models.Shop, error) {
	var shops []models.Shop
	if err := config.DB.Preload("Products").Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}
