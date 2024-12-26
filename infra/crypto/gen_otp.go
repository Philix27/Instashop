package crypto

import (
	"math/rand"
	"strconv"
)

// GenerateOTP generates a random 5-digit OTP
func GenerateOTP() string {
	randNum := 10000 + rand.Intn(90000) // Ensures it's a 5-digit number

	return strconv.Itoa(randNum)
}
