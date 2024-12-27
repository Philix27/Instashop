package crypto

// ValidateJWTToken validates a JWT token and checks if a specific claim matches the expected value
func ValidateAndCompareClaimToken(secretKey string, tokenString string, expectedValue string) bool {

	err, claimValue, _ := ValidateAndGetTokenPayload(secretKey, tokenString)
	println("Claim value:", claimValue)

	if err != nil {
		return false
	}

	return claimValue == expectedValue

}
