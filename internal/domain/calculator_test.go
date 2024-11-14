package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxes(t *testing.T) {
	// Arrange
	calculator := NewCalculator()
	operations := []Operation{
		{OperationType: BUY, Quantity: 10, UnitCost: 10},
		{OperationType: SELL, Quantity: 5, UnitCost: 20},
	}

	// Act
	taxes := calculator.CalculateTaxes(operations)

	// Assert
	assert.Equal(t, 2, len(taxes))
	assert.Equal(t, 0.0, taxes[0].Tax)
	assert.Equal(t, 0.0, taxes[1].Tax)
}
