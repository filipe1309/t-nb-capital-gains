package calculatecapitalgain_test

import (
	"bytes"
	ccg "capital-gains/internal/service/calculate-capital-gain"
	"path/filepath"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		caseName string
	}{
		{"case-0"},
		{"case-1"},
		{"case-2"},
		{"case-3"},
		{"case-4"},
		{"case-5"},
		{"case-6"},
		{"case-7"},
		{"case-8"},
		{"case-12"},
	}

	calculator := ccg.NewService()

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			t.Parallel()

			// Arrange
			inputPath := filepath.Join("testdata", tt.caseName+".input")
			input, err := os.Open(inputPath)
			if err != nil {
				t.Fatalf("failed to open input file: %v", err)
			}
			defer func(input *os.File) {
				err := input.Close()
				if err != nil {
					t.Fatalf("failed to close input file: %v", err)
				}
			}(input)

			outputPath := filepath.Join("testdata", tt.caseName+".golden")
			expected, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			// Act
			var output bytes.Buffer
			err = calculator.Calculate(input, &output);
			
			// Assert
			assert.Nil(t, err)
			assert.Equal(t, string(expected), output.String())
		})
	}
}
