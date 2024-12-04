package auth

import (
	"fmt"
	"log"
	"sample/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Get the secret key from the environment variable
func getSecretKey() []byte {
	secretKey, err := config.LoadSecret()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return []byte(secretKey)
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token with a given username and role
func GenerateJWT(username, role string) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // Set expiration to 2 hours
			Issuer:    "sample",
		},
	}
	secretKey := getSecretKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseJWT validates the token and returns the claims
func ParseJWT(tokenStr string) (*Claims, error) {
	secretKey := getSecretKey()
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
