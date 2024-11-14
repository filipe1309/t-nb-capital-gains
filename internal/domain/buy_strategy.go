package domain

import "capital-gains/internal/util"

type BuyStrategy struct{}

// Execute calculate tax for a buy operation
// new-weighted-average = ((current-shares * current-weighted-average) + (purchased-shares * purchase-value)) / (current-shares + purchased-shares)
func (s *BuyStrategy) Execute(c *Calculator, quantity int64, unitCost float64) *Tax {
	totalCurrentCost := float64(c.TotalSharesQuantity) * c.WeightedAveragePrice
	newCost := float64(quantity) * unitCost

	c.TotalSharesQuantity += quantity
	c.WeightedAveragePrice = util.RoundTwoDecimals((totalCurrentCost + newCost) / float64(c.TotalSharesQuantity))

	return &Tax{Tax: 0}
}
