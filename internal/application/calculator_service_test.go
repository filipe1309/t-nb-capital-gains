package application

import (
	"bytes"
	"capital-gains/internal/domain"
	"capital-gains/internal/infrastructure"
	"fmt"
	"path/filepath"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
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

	for _, tt := range tests {
		calculator := domain.NewCalculator()

		t.Run(tt.caseName, func(t *testing.T) {
			t.Parallel()

			// Arrange
			inputPath := filepath.Join("testdata", tt.caseName+".input")
			operationReader, err := infrastructure.NewFileOperationReader(inputPath)
			assert.Nil(t, err)
			defer operationReader.Close()
			var output bytes.Buffer
			taxesWriter := infrastructure.NewBufferTaxesWriter(&output)
			service := NewCalculatorService(calculator, operationReader, taxesWriter)
			outputPath := filepath.Join("testdata", tt.caseName+".golden")
			expected, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			// Act
			err = service.Calculate()

			// Assert
			assert.Nil(t, err)
			assert.Equal(t, string(expected), output.String())
		})
	}
}

type mockFileOperationReader struct {}

func (r *mockFileOperationReader) ReadOperations() ([][]domain.Operation, error) {
	return nil, fmt.Errorf("error reading operations")
}

func TestCalculateError(t *testing.T) {
	t.Run("case-9-no-file", func(t *testing.T) {
		// Arrange
		inputPath := filepath.Join("testdata", "case-9.input")
		_, err := infrastructure.NewFileOperationReader(inputPath)
		assert.Error(t, err)
	})

	t.Run("case-10-error-reading-json", func(t *testing.T) {
		// Arrange
		operationReader := &mockFileOperationReader{}
		var output bytes.Buffer
		taxesWriter := infrastructure.NewBufferTaxesWriter(&output)
		calculator := domain.NewCalculator()
		service := NewCalculatorService(calculator, operationReader, taxesWriter)

		// Act
		err := service.Calculate()

		// Assert
		assert.Error(t, err)
	})
}
