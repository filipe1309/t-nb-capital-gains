package infrastructure

import (
	"capital-gains/internal/domain"
	"encoding/json"
	"fmt"
	"os"
)

type STDOUTTaxesWriter struct{}

func NewSTDOUTTaxesWriter() *STDOUTTaxesWriter {
	return &STDOUTTaxesWriter{}
}

// WriteTaxes format taxes as JSON and output to stdout
func (w *STDOUTTaxesWriter) WriteTaxes(taxes domain.Taxes) error {
	// Format taxes as JSON and output to stdout
	if taxesJSON, err := json.Marshal(taxes); err != nil {
		return fmt.Errorf("error marshalling taxes: %v", err)
	} else {
		fmt.Fprintln(os.Stdout, string(taxesJSON))
	}

	return nil
}
