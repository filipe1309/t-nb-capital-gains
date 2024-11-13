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

// func ReadOperationsFromJSON(reader io.Reader) ([][]Operation, error) {
// 	var operations [][]Operation
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		if scanner.Text() == "" {
// 			break
// 		}

// 		var currentOperations []Operation
// 		if err := json.Unmarshal([]byte(line), &currentOperations); err != nil {
// 			log.Printf("error unmarshalling operations: %v", err)
// 			return nil, err
// 		}

// 		operations = append(operations, currentOperations)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	return operations, nil
// }
