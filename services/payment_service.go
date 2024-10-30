// services/payment_service.go
package services

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/joshua468/ecommerce/config"
	"github.com/joshua468/ecommerce/models"
)

type PaymentResponse struct {
	PaymentURL string `json:"payment_url"`
}

// Update the argument to use models.PaymentRequest
func ProcessPayment(paymentRequest models.PaymentRequest) (PaymentResponse, error) {
	body, _ := json.Marshal(paymentRequest)
	req, err := http.NewRequest("POST", config.FlutterwaveURL, bytes.NewBuffer(body))
	if err != nil {
		return PaymentResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+config.FlutterwaveKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	var paymentResponse PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return PaymentResponse{}, err
	}
	return paymentResponse, nil
}
