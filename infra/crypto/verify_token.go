package crypto

// ValidateJWTToken validates a JWT token and checks if a specific claim matches the expected value
func ValidateToken(tokenString string, expectedValue string) bool {

	err, claimValue := GetTokenPayload(tokenString)
	println("Claim value:", claimValue)
	
	if err != nil {
		return false
	}

	return claimValue == expectedValue

}
