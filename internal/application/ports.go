package application

import (
	"capital-gains/internal/domain"
)

// OperationReader is an interface that defines the ReadOperations method.
type OperationReader interface {
	ReadOperations() ([][]domain.Operation, error)
}

// TaxesWriter is an interface that defines the WriteTaxes method.
type TaxesWriter interface {
	WriteTaxes(domain.Taxes) error
}
