package infrastructure

import (
	"capital-gains/internal/domain"

	"net/http"
)

type APIOperationReader struct {
	client *http.Client
}

func NewAPIOperationReader(client *http.Client) *APIOperationReader {
	return &APIOperationReader{client: client}
}

func (r *APIOperationReader) ReadOperations() ([][]domain.Operation, error) {
	var operations [][]domain.Operation

	// TODO: Implement the ReadOperations method for the APIOperationReader

	return operations, nil
}
