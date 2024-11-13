package domain

// OperationType represents the type of operation
type OperationType string

const (
	BUY  OperationType = "buy"
	SELL OperationType = "sell"
)

type Operation struct {
	OperationType OperationType `json:"operation"`
	UnitCost      float64       `json:"unit-cost"`
	Quantity      int64         `json:"quantity"`
}

// func ReadOperationsFromJSON() ([]Operation, error) {
// 	return []Operation{
// 		{OperationType: BUY, UnitCost: 10.00, Quantity: 100},
// 		{OperationType: SELL, UnitCost: 15.00, Quantity: 50},
// 		{OperationType: SELL, UnitCost: 15.00, Quantity: 50},
// 	}
// }
