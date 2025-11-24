package random

import (
	"crypto/rand"
	"math/big"
)

// OTP generates a cryptographically secure one-time password (OTP).
// The default length is 6 digits if no length is specified.
// Only numeric characters (0-9) are used.
// Returns an error if the cryptographic random number generator fails.
//
// Length handling:
//   - No length or length <= 0: defaults to 6 digits
//   - Length > 64: silently clamped to 64 digits (maximum allowed)
//
// Example:
//
//	otp, err := random.OTP()      // generates 6-digit OTP
//	if err != nil {
//	    return err
//	}
//	otp, err := random.OTP(8)     // generates 8-digit OTP
//	if err != nil {
//	    return err
//	}
func OTP(length ...int) (string, error) {
	// Default to 6 digits
	otpLength := 6
	if len(length) > 0 && length[0] > 0 {
		otpLength = length[0]
	}

	// Prevent excessive memory allocation (max 64 digits)
	const maxOTPLength = 64
	if otpLength > maxOTPLength {
		otpLength = maxOTPLength
	}

	const digits = "0123456789"
	maxIndex := big.NewInt(int64(len(digits)))

	result := make([]byte, otpLength)
	for i := 0; i < otpLength; i++ {
		n, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}
		result[i] = digits[n.Int64()]
	}

	return string(result), nil
}
