package random

import (
	"regexp"
	"testing"
)

func TestOTP(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			var otp string
			var err error
			if tt.length == nil {
				otp, err = OTP()
			} else {
				otp, err = OTP(tt.length[0])
			}

			if err != nil {
				t.Fatalf("OTP() unexpected error: %v", err)
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
	t.Parallel()

	// Generate multiple OTPs and ensure they're not all the same
	otps := make(map[string]bool)
	for i := 0; i < 100; i++ {
		otp, err := OTP()
		if err != nil {
			t.Fatalf("OTP() unexpected error: %v", err)
		}
		otps[otp] = true
	}

	// With 100 6-digit OTPs, we should have many unique values
	// (probability of collision is extremely low)
	if len(otps) < 95 {
		t.Errorf("Expected at least 95 unique OTPs out of 100, got %d", len(otps))
	}
}

func TestOTP_ZeroLength(t *testing.T) {
	t.Parallel()

	otp, err := OTP(0)
	if err != nil {
		t.Fatalf("OTP(0) unexpected error: %v", err)
	}

	// Zero length should default to 6
	if len(otp) != 6 {
		t.Errorf("OTP(0) length = %v, want 6", len(otp))
	}
}

func TestOTP_NegativeLength(t *testing.T) {
	t.Parallel()

	otp, err := OTP(-5)
	if err != nil {
		t.Fatalf("OTP(-5) unexpected error: %v", err)
	}

	// Negative length should default to 6
	if len(otp) != 6 {
		t.Errorf("OTP(-5) length = %v, want 6", len(otp))
	}
}
