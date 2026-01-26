package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretString = []byte("supersecret")

func GenerateToken(userID string, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my-app",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(SecretString)
	return tokenStr, err
}
