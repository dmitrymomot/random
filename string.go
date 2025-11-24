package random

import (
	"math/rand"
	"strings"
)

const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
)

// String generates a random string of the specified length using the provided character sets.
// If no character sets are provided, it defaults to Alphanumeric.
// This function uses math/rand (automatically seeded in Go 1.20+) and is NOT cryptographically secure.
// Use OTP() for security-sensitive operations.
//
// Example:
//
//	random.String(16)                          // alphanumeric string
//	random.String(32, random.Hex)              // hex string
//	random.String(20, random.Uppercase, random.Numeric) // uppercase + digits
func String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
