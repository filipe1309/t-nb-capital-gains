package calculatecapitalgain_test

import (
	"bytes"
	calculatecapitalgain2 "capital-gains/internal/service/calculate-capital-gain"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		caseName string
	}{
		{"case-1"},
		{"case-2"},
		{"case-12"},
		{"case-3"},
		{"case-4"},
		{"case-5"},
		{"case-6"},
		{"case-7"},
		{"case-8"},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			t.Parallel()

			input, err := os.Open("stubs/" + tt.caseName + "-input.txt")
			if err != nil {
				t.Fatalf("failed to open input file: %v", err)
			}
			defer func(input *os.File) {
				err := input.Close()
				if err != nil {
					t.Fatalf("failed to close input file: %v", err)
				}
			}(input)

			expected, err := os.ReadFile("stubs/" + tt.caseName + "-output.txt")
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			var output bytes.Buffer
			if err := calculatecapitalgain2.Calculate2(input, &output); err != nil {
				t.Fatalf("failed to run: %v", err)
			}

			assert.Equal(t, string(expected), output.String())
		})
	}
}
