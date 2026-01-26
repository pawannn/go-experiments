package main

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
