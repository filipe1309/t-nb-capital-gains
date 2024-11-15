package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxes(t *testing.T) {
	tests := []struct {
		caseName string
		operations []Operation
		verify func(t *testing.T, taxes Taxes)
	}{
		{
			caseName: "case-7",
			operations: []Operation{
				{OperationType: BUY, Quantity: 10000, UnitCost: 10},
				{OperationType: SELL, Quantity: 5000, UnitCost: 2},
				{OperationType: SELL, Quantity: 2000, UnitCost: 20},
				{OperationType: SELL, Quantity: 2000, UnitCost: 20},
				{OperationType: SELL, Quantity: 1000, UnitCost: 25},
				{OperationType: BUY, Quantity: 10000, UnitCost: 20},
				{OperationType: SELL, Quantity: 5000, UnitCost: 15},
				{OperationType: SELL, Quantity: 4350, UnitCost: 30},
				{OperationType: SELL, Quantity: 650, UnitCost: 30},
			},
			verify: func(t *testing.T, taxes Taxes) {
				assert.Equal(t, 9, len(taxes))
				assert.Equal(t, 0.0, taxes[0].Tax)
				assert.Equal(t, 0.0, taxes[1].Tax)
				assert.Equal(t, 0.0, taxes[2].Tax)
				assert.Equal(t, 0.0, taxes[3].Tax)
				assert.Equal(t, 3000.0, taxes[4].Tax)
				assert.Equal(t, 0.0, taxes[5].Tax)
				assert.Equal(t, 0.0, taxes[6].Tax)
				assert.Equal(t, 3700.0, taxes[7].Tax)
				assert.Equal(t, 0.0, taxes[8].Tax)
			},
		},
	}

	for _, tt := range tests {
		// Arrange
		calculator := NewCalculator()

		// Act
		taxes := calculator.CalculateTaxes(tt.operations)

		// Assert
		tt.verify(t, taxes)
	}
}

func TestCalculateTaxesWrongOperation(t *testing.T) {
	// Arrange
	calculator := NewCalculator()
	operations := []Operation{
		{OperationType: "UNKNOWN", Quantity: 10, UnitCost: 10},
	}

	// Act
	taxes := calculator.CalculateTaxes(operations)

	// Assert
	assert.Equal(t, 1, len(taxes))
	assert.Equal(t, taxes[0], Tax{Tax: 0})
}
