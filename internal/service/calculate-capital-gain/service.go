package calculatecapitalgain

import (
	"capital-gains/internal/domain"
	"encoding/json"
	"fmt"
)

type CalculatorService struct{}

// NewService creates a new instance of the CalculatorService
func NewService() *CalculatorService {
	return &CalculatorService{}
}

// Calculate reads operations from stdin, calculates taxes and outputs the result to stdout
func (srv *CalculatorService) Calculate(reader Reader, writer Writer) error {
	multiOperations, err := domain.ReadOperationsFromJSON(reader)
	if err != nil {
		return fmt.Errorf("error reading operations: %v", err)
	}

	for _, operations := range multiOperations {

		// Process operations and calculate taxes
		calculator := domain.NewCalculator()
		taxes := calculator.CalculateTaxes(operations)

		// Format taxes as JSON and output to stdout
		if taxesJSON, err := json.Marshal(taxes); err != nil {
			return fmt.Errorf("error marshalling taxes: %v", err)
		} else {
			fmt.Fprintln(writer, string(taxesJSON))
		}
	}

	return nil
}
