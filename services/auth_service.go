package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joshua468/ecommerce/config"
	"github.com/joshua468/ecommerce/models"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key")

func RegisterUser(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(username, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}
	token, err := generateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprint(userID),
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
