package random

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("default alphanumeric", func(t *testing.T) {
		t.Parallel()

		length := uint8(10)
		result := String(length)

		require.Equal(t, int(length), len(result))

		// Verify all characters are alphanumeric
		alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
		require.True(t, alphanumericRegex.MatchString(result))
	})

	t.Run("uppercase only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Uppercase)

		require.Equal(t, int(length), len(result))

		// Verify all characters are uppercase
		uppercaseRegex := regexp.MustCompile(`^[A-Z]+$`)
		require.True(t, uppercaseRegex.MatchString(result))
	})

	t.Run("lowercase only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Lowercase)

		require.Equal(t, int(length), len(result))

		// Verify all characters are lowercase
		lowercaseRegex := regexp.MustCompile(`^[a-z]+$`)
		require.True(t, lowercaseRegex.MatchString(result))
	})

	t.Run("alphabetic only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Alphabetic)

		require.Equal(t, int(length), len(result))

		// Verify all characters are alphabetic
		alphabeticRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
		require.True(t, alphabeticRegex.MatchString(result))
	})

	t.Run("numeric only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Numeric)

		require.Equal(t, int(length), len(result))

		// Verify all characters are numeric
		numericRegex := regexp.MustCompile(`^[0-9]+$`)
		require.True(t, numericRegex.MatchString(result))
	})

	t.Run("alphanumeric explicit", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Alphanumeric)

		require.Equal(t, int(length), len(result))

		// Verify all characters are alphanumeric
		alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
		require.True(t, alphanumericRegex.MatchString(result))
	})

	t.Run("hex only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Hex)

		require.Equal(t, int(length), len(result))

		// Verify all characters are hex
		hexRegex := regexp.MustCompile(`^[0-9a-f]+$`)
		require.True(t, hexRegex.MatchString(result))
	})

	t.Run("symbols only", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, Symbols)

		require.Equal(t, int(length), len(result))

		// Verify all characters are from the symbols set
		for _, char := range result {
			require.True(t, strings.ContainsRune(Symbols, char))
		}
	})

	t.Run("multiple charsets combined", func(t *testing.T) {
		t.Parallel()

		length := uint8(30)
		result := String(length, Uppercase, Numeric)

		require.Equal(t, int(length), len(result))

		// Verify all characters are from combined charset
		combinedRegex := regexp.MustCompile(`^[A-Z0-9]+$`)
		require.True(t, combinedRegex.MatchString(result))
	})

	t.Run("all charsets combined", func(t *testing.T) {
		t.Parallel()

		length := uint8(50)
		result := String(length, Alphanumeric, Symbols)

		require.Equal(t, int(length), len(result))

		// Verify all characters are from the combined charset
		for _, char := range result {
			require.True(t,
				strings.ContainsRune(Alphanumeric, char) ||
					strings.ContainsRune(Symbols, char),
			)
		}
	})

	t.Run("zero length", func(t *testing.T) {
		t.Parallel()

		result := String(0)

		require.Equal(t, 0, len(result))
		require.Equal(t, "", result)
	})

	t.Run("single character", func(t *testing.T) {
		t.Parallel()

		result := String(1)

		require.Equal(t, 1, len(result))

		// Verify it's alphanumeric
		alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]$`)
		require.True(t, alphanumericRegex.MatchString(result))
	})

	t.Run("randomness check", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		results := make(map[string]bool)

		// Generate multiple strings
		for i := 0; i < 100; i++ {
			result := String(length)
			results[result] = true
		}

		// Should have many unique strings
		require.Greater(t, len(results), 95, "Expected at least 95 unique strings out of 100")
	})

	t.Run("custom charset", func(t *testing.T) {
		t.Parallel()

		customCharset := "ABC123"
		length := uint8(20)
		result := String(length, customCharset)

		require.Equal(t, int(length), len(result))

		// Verify all characters are from custom charset
		for _, char := range result {
			require.True(t, strings.ContainsRune(customCharset, char))
		}
	})

	t.Run("empty charset defaults to alphanumeric", func(t *testing.T) {
		t.Parallel()

		length := uint8(20)
		result := String(length, "")

		require.Equal(t, int(length), len(result))

		// Should default to alphanumeric
		alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
		require.True(t, alphanumericRegex.MatchString(result))
	})

	t.Run("max uint8 length", func(t *testing.T) {
		t.Parallel()

		length := uint8(255)
		result := String(length)

		require.Equal(t, int(length), len(result))
	})
}

func TestString_CharacterDistribution(t *testing.T) {
	t.Parallel()

	t.Run("uppercase and lowercase distribution", func(t *testing.T) {
		t.Parallel()

		length := uint8(100)
		uppercaseCount := 0
		lowercaseCount := 0

		for i := 0; i < 10; i++ {
			result := String(length, Alphabetic)
			for _, char := range result {
				if char >= 'A' && char <= 'Z' {
					uppercaseCount++
				} else if char >= 'a' && char <= 'z' {
					lowercaseCount++
				}
			}
		}

		// Both should appear with reasonable frequency
		totalChars := uppercaseCount + lowercaseCount
		require.Equal(t, 1000, totalChars)

		// Each should have at least 30% representation
		require.Greater(t, uppercaseCount, 300)
		require.Greater(t, lowercaseCount, 300)
	})

	t.Run("alphanumeric distribution", func(t *testing.T) {
		t.Parallel()

		length := uint8(100)
		letterCount := 0
		digitCount := 0

		for i := 0; i < 10; i++ {
			result := String(length, Alphanumeric)
			for _, char := range result {
				if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
					letterCount++
				} else if char >= '0' && char <= '9' {
					digitCount++
				}
			}
		}

		// Both letters and digits should appear
		totalChars := letterCount + digitCount
		require.Equal(t, 1000, totalChars)

		// Letters should be more common due to larger charset (52 vs 10)
		require.Greater(t, letterCount, digitCount)
	})
}

func TestString_Constants(t *testing.T) {
	t.Parallel()

	t.Run("verify constant values", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Uppercase)
		assert.Equal(t, "abcdefghijklmnopqrstuvwxyz", Lowercase)
		assert.Equal(t, Uppercase+Lowercase, Alphabetic)
		assert.Equal(t, "0123456789", Numeric)
		assert.Equal(t, Alphabetic+Numeric, Alphanumeric)
		assert.Equal(t, "`~!@#$%^&*()-_+={}[]|\\;:\"<>,./?", Symbols)
		assert.Equal(t, Numeric+"abcdef", Hex)
	})

	t.Run("verify constant lengths", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, 26, len(Uppercase))
		assert.Equal(t, 26, len(Lowercase))
		assert.Equal(t, 52, len(Alphabetic))
		assert.Equal(t, 10, len(Numeric))
		assert.Equal(t, 62, len(Alphanumeric))
		assert.Equal(t, 16, len(Hex))
		assert.Equal(t, 31, len(Symbols))
	})
}
