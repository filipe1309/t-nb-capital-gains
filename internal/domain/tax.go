package domain

import "fmt"

// Rate is the tax rate
const Rate = 0.20

// FreeTaxValue is the limit for no tax
const FreeTaxValue = 20000

// Tax represents the tax amount
type Tax struct {
	Tax float64 `json:"tax"`
}

// Taxes is a list of Tax
type Taxes []Tax

// MarshalJSON is called by json.Marshal to encode the Tax as JSON
func (t Tax) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"tax":%.2f}`, t.Tax)), nil
}
