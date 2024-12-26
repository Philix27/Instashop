package crypto

// ValidateJWTToken validates a JWT token and checks if a specific claim matches the expected value
func ValidateAndCompareClaimToken(tokenString string, expectedValue string) bool {

	err, claimValue := ValidateAndGetTokenPayload(tokenString)
	println("Claim value:", claimValue)

	if err != nil {
		return false
	}

	return claimValue == expectedValue

}
