package domain

// SellStrategy is a struct that implements the OperationStrategy interface
type SellStrategy struct{}

// Execute calculate tax for a sell operation
func (s *SellStrategy) Execute(c *Calculator, quantity int64, unitCost float64) *Tax {
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
