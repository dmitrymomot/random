// Package random provides utilities for generating random strings, one-time passwords,
// and performing weighted random selections.
//
// # Random String Generation
//
// The [String] function generates random strings using math/rand. It is fast and suitable
// for non-security-sensitive purposes such as generating IDs, test data, or session tokens
// where cryptographic security is not required.
//
// Predefined character set constants are available for common use cases:
// Uppercase, Lowercase, Alphabetic, Numeric, Alphanumeric, Symbols, and Hex.
//
//	// Generate a random alphanumeric string of length 16
//	randomID := random.String(16)
//
//	// Generate a random hex string of length 32
//	randomHex := random.String(32, random.Hex)
//
//	// Combine multiple character sets
//	randomStr := random.String(20, random.Uppercase, random.Numeric)
//
// # Cryptographically Secure OTP Generation
//
// The [OTP] function generates cryptographically secure one-time passwords using crypto/rand.
// It is suitable for security-sensitive operations such as generating authentication codes,
// password reset tokens, and multi-factor authentication (MFA) codes.
//
// The default OTP length is 6 digits. Custom lengths can be specified.
// Only numeric characters (0-9) are used in OTP generation.
//
//	// Generate a 6-digit OTP (default)
//	otp, err := random.OTP()
//	if err != nil {
//	    return err
//	}
//
//	// Generate an 8-digit OTP
//	otp, err := random.OTP(8)
//	if err != nil {
//	    return err
//	}
//
// # Weighted Random Selection
//
// The package provides several functions for performing weighted random selection from
// different data structures. These use math/rand and are suitable for non-security-sensitive
// probabilistic selections.
//
// [GetRandomWithProbabilities] selects a random item from a slice with custom probabilities.
// Probabilities are relative weights and do not need to sum to any specific value.
//
//	items := []any{"apple", "banana", "cherry"}
//	probabilities := []float64{0.5, 0.3, 0.2}
//	selected := random.GetRandomWithProbabilities(items, probabilities)
//
// [GetRandomStructWithProbabilities] selects a random item from a slice of structures
// that implement the GetProbability() float64 interface method.
//
//	type Option struct {
//	    Name string
//	}
//
//	func (o Option) GetProbability() float64 {
//	    // return probability weight
//	    return 0.5
//	}
//
//	options := []interface{ GetProbability() float64 }{/* ... */}
//	selected := random.GetRandomStructWithProbabilities(options)
//
// [GetRandomMapItemWithProbabilities] selects a random key from a map where values
// are probability weights.
//
//	items := map[string]float64{
//	    "common":   0.7,
//	    "uncommon": 0.2,
//	    "rare":     0.1,
//	}
//	selected := random.GetRandomMapItemWithProbabilities(items)
//
// [GetRandomMapItemWithPercent] selects a random key from a map where values are
// percentages (0-100). This function attempts random selection up to 100 times
// before returning the first available item.
//
//	drops := map[string]float64{
//	    "common":   50.0,
//	    "uncommon": 30.0,
//	    "rare":     20.0,
//	}
//	selected := random.GetRandomMapItemWithPercent(drops)
//
// # Security Guidance
//
// WARNING: Use the appropriate function for your security requirements:
//
//   - [String] uses math/rand and is NOT cryptographically secure.
//     Use only for non-security purposes (IDs, test data, display strings).
//
//   - [OTP] uses crypto/rand and IS cryptographically secure.
//     Use for security-sensitive operations (passwords, tokens, MFA codes, session IDs).
//
// All probability-based selection functions use math/rand and are NOT suitable for
// security-sensitive applications. Use them only for game mechanics, simulations,
// loot tables, and other non-security purposes.
//
// # Error Handling
//
// The [OTP] function returns an error if the cryptographic random number generator fails.
// All other functions return sensible defaults (nil, empty string) on invalid input rather
// than panicking.
package random
