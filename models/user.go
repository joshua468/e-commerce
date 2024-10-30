package models

import "time"

type User struct {
	ID  uint 			`json:"id" gorm:"primaryKey"`
	Username string     `json:"username" gorm:"unique"`
	Password string		`json:"-"`
	Email string		`json:"email" gorm:"unique"`
	Role  string		`josn:"role"`
	Balance float64		`json:"balance"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}