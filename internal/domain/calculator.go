package domain

// Calculator represents the operations values to calculate taxes
type Calculator struct {
	strategies           map[OperationType]OperationStrategy
	TotalLoss            float64
	TotalSharesQuantity  int64
	WeightedAveragePrice float64
}

// NewCalculator creates a new calculator
func NewCalculator() *Calculator {
	return &Calculator{
		strategies: map[OperationType]OperationStrategy{
			BUY:  new(BuyStrategy),
			SELL: new(SellStrategy),
		},
	}
}

// CalculateTaxes calculate taxes for a list of operations
func (c *Calculator) CalculateTaxes(operations []Operation) Taxes {
	taxes := make(Taxes, len(operations))
	for i, operation := range operations {
		taxes[i] = *c.calculateOperationTax(operation)
	}
	return taxes
}

// calculateOperationTax calculate tax for a single operation
func (c *Calculator) calculateOperationTax(operation Operation) *Tax {
	strategy, exists := c.strategies[operation.OperationType]
	if !exists {
		return nil
	}
	return strategy.Execute(c, operation.Quantity, operation.UnitCost)
}
