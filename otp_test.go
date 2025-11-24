package random

import (
	"regexp"
	"testing"
)

func TestOTP(t *testing.T) {
	tests := []struct {
		name   string
		length []int
		want   int
	}{
		{"default 6 digits", nil, 6},
		{"custom 8 digits", []int{8}, 8},
		{"custom 4 digits", []int{4}, 4},
		{"custom 10 digits", []int{10}, 10},
	}

	numericRegex := regexp.MustCompile(`^[0-9]+$`)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var otp string
			if tt.length == nil {
				otp = OTP()
			} else {
				otp = OTP(tt.length[0])
			}

			// Check length
			if got := len(otp); got != tt.want {
				t.Errorf("OTP() length = %v, want %v", got, tt.want)
			}

			// Check if all characters are numeric
			if !numericRegex.MatchString(otp) {
				t.Errorf("OTP() = %v, contains non-numeric characters", otp)
			}
		})
	}
}

func TestOTP_Randomness(t *testing.T) {
	// Generate multiple OTPs and ensure they're not all the same
	otps := make(map[string]bool)
	for i := 0; i < 100; i++ {
		otp := OTP()
		otps[otp] = true
	}

	// With 100 6-digit OTPs, we should have many unique values
	// (probability of collision is extremely low)
	if len(otps) < 95 {
		t.Errorf("Expected at least 95 unique OTPs out of 100, got %d", len(otps))
	}
}
