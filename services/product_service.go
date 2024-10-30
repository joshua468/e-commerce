package services

import(
"github.com/joshua468/ecommerce/models"
"github.com/joshua468/ecommerce/config"
)

func CreateProduct(product *models.Product) error{
	if err:= config.DB.Create(product).Error; err!=nil {
		return err
	}
	return nil
}

func GetAvailableProducts() ([]models.Product,error) {
var products []models.Product
if err:= config.DB.Find(&products).Error;err!=nil {
	return nil,err
}
return products,nil
}

