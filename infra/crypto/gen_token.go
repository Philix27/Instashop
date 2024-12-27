package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CreateJWTToken creates a JWT token with the OTP as payload
func CreateJWTToken(secretKey []byte, payload string, role string, d time.Duration) (tokenString string, err error) {
	// Define the claims
	claims := jwt.MapClaims{
		"payload": payload,
		"role":    role,
		"exp":     time.Now().Add(d).Unix(), // Token expiration in 10 minutes
		"iat":     time.Now().Unix(),        // Issued at time
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		return
	}
	return
}
