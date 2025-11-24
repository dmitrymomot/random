package random

import (
	"crypto/rand"
	"math/big"
)

// OTP generates a cryptographically secure one-time password (OTP).
// The default length is 6 digits if no length is specified.
// Only numeric characters (0-9) are used.
//
// Example:
//
//	otp := random.OTP()      // generates 6-digit OTP
//	otp := random.OTP(8)     // generates 8-digit OTP
func OTP(length ...int) string {
	// Default to 6 digits
	otpLength := 6
	if len(length) > 0 && length[0] > 0 {
		otpLength = length[0]
	}

	const digits = "0123456789"
	maxIndex := big.NewInt(int64(len(digits)))

	result := make([]byte, otpLength)
	for i := 0; i < otpLength; i++ {
		n, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			panic(err)
		}
		result[i] = digits[n.Int64()]
	}

	return string(result)
}
