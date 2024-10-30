package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	OwnerID     uint      `json:"owner_id" gorm:"not null"`          // Foreign key to User
	Products    []Product `json:"products" gorm:"foreignKey:ShopID"` // One-to-many relationship with Product
}
