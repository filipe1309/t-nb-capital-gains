package util

import (
	"fmt"
	"testing"
)

func TestRoundTwoDecimals(t *testing.T) {
	// Test cases
	tests := []struct {
		value    float64
		expected float64
	}{
		{1.123, 1.12},
		{1.126, 1.13},
		{1.125, 1.13},
	}

	// Test
	for k, test := range tests {
		t.Run(fmt.Sprint(k), func(t *testing.T) {
			result := RoundTwoDecimals(test.value)
			if result != test.expected {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		})
	}
}
