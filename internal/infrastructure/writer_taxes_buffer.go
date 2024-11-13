package infrastructure

import (
	"bytes"
	"capital-gains/internal/domain"
	"encoding/json"
	"fmt"
)

type BufferTaxesWriter struct{
	buffer *bytes.Buffer
}

func NewBufferTaxesWriter(buffer *bytes.Buffer) *BufferTaxesWriter {
	return &BufferTaxesWriter{buffer: buffer}
}

// WriteTaxes writes taxes to a buffer
func (w *BufferTaxesWriter) WriteTaxes(taxes domain.Taxes) error {
	if taxesJSON, err := json.Marshal(taxes); err != nil {
		return fmt.Errorf("error marshalling taxes: %v", err)
	} else {
		fmt.Fprintln(w.buffer, string(taxesJSON))
	}

	return nil
}
