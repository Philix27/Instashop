package utils

import "regexp"

// ValidateEmail checks if the given email is valid
func ValidateEmail(email string) bool {
	// Define a regex pattern for a basic email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex
	re := regexp.MustCompile(emailRegex)
	// Match the input email against the regex
	return re.MatchString(email)
}
