package main

import "testing"

// Test cases for ValidateCreditCard
func TestValidateCreditCard(t *testing.T) {
	testCases := []struct {
		description string // Adding a description for clarity
		cardNum     string
		expected    bool
	}{
		{"valid card", "4532015112830366", true},                        // Valid card
		{"invalid card - wrong check digit", "4532015112830361", false}, // Invalid card
		{"empty input", "", false},                                      // Empty input
		{"non-numeric input", "abcdefg", false},                         // Non-numeric input
		{"numeric but invalid", "1234567890123456", false},              // Numeric but invalid
	}

	for _, tc := range testCases {
		result := ValidateCreditCard(tc.cardNum)
		if result != tc.expected {
			t.Errorf("Failed %s: ValidateCreditCard(%s) = %t, expected %t", tc.description, tc.cardNum, result, tc.expected)
		}
	}
}
