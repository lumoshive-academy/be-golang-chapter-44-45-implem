package main

import (
	"be-golang-chapter-44-45-implem/config"
	"be-golang-chapter-44-45-implem/helper"
	"fmt"
)

func init() {
	config.ReadConfig()
}

func main() {
	// Data user
	userID := "123"
	email := "user@example.com"
	username := "user123"

	// Secret key untuk HMAC
	secretKey := config.Configs.Secret

	// Generate token
	token := helper.GenerateToken(userID, email, username, secretKey)
	fmt.Println("Generated Token:", token)

	// Validate token
	isValid, message := helper.ValidasialidateToken(token, secretKey)
	if isValid {
		fmt.Println("Token is valid! Data:", message)
	} else {
		fmt.Println("Token is invalid:", message)
	}
}
