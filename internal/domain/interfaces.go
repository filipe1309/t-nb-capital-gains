package domain

// OperationReader is an interface that defines the ReadOperations method.
type OperationReader interface {
	ReadOperations() ([][]Operation, error)
}

// TaxesWriter is an interface that defines the WriteTaxes method.
type TaxesWriter interface {
	WriteTaxes(Taxes) error
}
