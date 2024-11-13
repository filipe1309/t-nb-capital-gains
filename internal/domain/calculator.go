package domain

import (
	"capital-gains/internal/util"
)

// Calculator represents the operations values to calculate taxes
type Calculator struct {
	TotalLoss            float64
	TotalSharesQuantity  int64
	WeightedAveragePrice float64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

// CalculateTaxes calculate taxes for a list of operations
func (c *Calculator) CalculateTaxes(operations []Operation) Taxes {
	taxes := make(Taxes, len(operations))
	for i, operation := range operations {
		tax := c.calculateTax(operation)
		if tax != nil {
			taxes[i] = *tax
		}
	}
	return taxes
}

// calculateTax calculate tax for a single operation
func (c *Calculator) calculateTax(operation Operation) *Tax {
	switch operation.OperationType {
	case BUY:
		return c.buy(operation.Quantity, operation.UnitCost)
	case SELL:
		return c.sell(operation.Quantity, operation.UnitCost)
	default:
		return nil
	}
}

// buy calculate tax for a buy operation
// new-weighted-average = ((current-shares * current-weighted-average) + (purchased-shares * purchase-value)) / (current-shares + purchased-shares)
func (c *Calculator) buy(quantity int64, unitCost float64) *Tax {
	totalCurrentCost := float64(c.TotalSharesQuantity) * c.WeightedAveragePrice
	newCost := float64(quantity) * unitCost

	c.TotalSharesQuantity += quantity
	c.WeightedAveragePrice = util.RoundTwoDecimals((totalCurrentCost + newCost) / float64(c.TotalSharesQuantity))

	return &Tax{Tax: 0}
}

// sell calculate tax for a sell operation
func (c *Calculator) sell(quantity int64, unitCost float64) *Tax {
	profit := (unitCost - c.WeightedAveragePrice) * float64(quantity)

	if unitCost*float64(quantity) <= FreeTaxValue {
		c.handleFreeTax(quantity, profit)
		return &Tax{Tax: 0}
	}

	if profit <= 0 {
		c.handleLossSale(quantity, profit)
		return &Tax{Tax: 0}
	}

	return c.handleProfitSale(quantity, profit)
}

// handleFreeTax handle free tax for a sell operation
func (c *Calculator) handleFreeTax(quantity int64, profit float64) {
	if profit <= 0 {
		c.TotalLoss += -profit
	}
	c.TotalSharesQuantity -= quantity
}

// handleLossSale handle loss sale for a sell operation
func (c *Calculator) handleLossSale(quantity int64, profit float64) {
	c.TotalSharesQuantity -= quantity
	c.TotalLoss += -profit
}

// handleProfitSale handle profit sale for a sell operation
func (c *Calculator) handleProfitSale(quantity int64, profit float64) *Tax {
	if c.TotalLoss > 0 {
		if c.TotalLoss >= profit {
			c.TotalLoss -= profit
			c.TotalSharesQuantity -= quantity
			return &Tax{Tax: 0}
		}

		profit -= c.TotalLoss
		c.TotalLoss = 0
	}

	c.TotalSharesQuantity -= quantity

	return &Tax{Tax: profit * Rate}
}
