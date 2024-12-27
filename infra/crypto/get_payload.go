package crypto

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ValidateJWTToken validates a JWT token and checks if a specific claim matches the expected value
func ValidateAndGetTokenPayload(secretKey string, tokenString string) (err error, payload string, role string) {

	// Extract and validate claims
	err, claims := validate(secretKey, tokenString)

	if err != nil {
		return errors.New("invalid token"), "", ""
	}
	// Check if the expected claim exists and matches the expected value
	payload, ok := claims["payload"].(string)
	if !ok {
		return fmt.Errorf("No payload"), "", ""
	}

	role, ok = claims["role"].(string)
	if !ok {
		return fmt.Errorf("No role provided"), "", ""
	}

	return
}

func validate(secretKey string, tokenString string) (err error, claim jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what you expect (HMAC-SHA256 in this case)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return errors.New("invalid token"), nil
	}

	// Check if the token is valid
	if !token.Valid {
		return fmt.Errorf("token is not valid"), nil
	}

	// Extract and validate claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims"), nil
	}

	// Validate expiration (optional, as it's included in the library's checks)
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return fmt.Errorf("token has expired"), nil
		}
	}

	return nil, claims
}
