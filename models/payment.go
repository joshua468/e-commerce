// models/payment_models.go
package models

// PaymentRequest struct to hold payment request details
type PaymentRequest struct {
	OrderID  uint    `json:"order_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// Payment struct to represent payment information
type Payment struct {
	ID         uint    `json:"id"`
	OrderID    uint    `json:"order_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
	PaymentURL string  `json:"payment_url"`
}
