package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	// Arrange
	tax := Tax{Tax: 10}

	// Act
	json, err := tax.MarshalJSON()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, `{"tax":10.00}`, string(json))
}
