package domain

type OperationStrategy interface {
	Execute(c *Calculator, quantity int64, unitCost float64) *Tax
}
