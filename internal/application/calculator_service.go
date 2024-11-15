package application

import (
	"capital-gains/internal/domain"
	"fmt"
)

// CalculatorService is a struct that contains a calculator, an operation reader and a taxes writer
type CalculatorService struct {
	calculator      *domain.Calculator
	operationReader domain.OperationReader
	taxesWriter     domain.TaxesWriter
}

// NewCalculatorService creates a new instance of the CalculatorService
func NewCalculatorService(
	calculator *domain.Calculator,
	operationReader domain.OperationReader,
	taxesWriter domain.TaxesWriter,
) *CalculatorService {
	return &CalculatorService{
		calculator:      calculator,
		operationReader: operationReader,
		taxesWriter:     taxesWriter,
	}
}

// Calculate reads operations, calculates taxes and writes them to the output
func (srv *CalculatorService) Calculate() error {
	multiOperations, err := srv.operationReader.ReadOperations()
	if err != nil {
		return fmt.Errorf("error reading operations: %v", err)
	}

	for _, operations := range multiOperations {
		taxes := srv.calculator.CalculateTaxes(operations)
		srv.taxesWriter.WriteTaxes(taxes)
	}

	return nil
}
