package domain

// OperationStrategy is an interface that defines the Execute method.
type OperationStrategy interface {
	Execute(c *Calculator, quantity int64, unitCost float64) *Tax
}
