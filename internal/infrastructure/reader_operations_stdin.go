package infrastructure

import (
	"bufio"
	"capital-gains/internal/domain"
	"encoding/json"
	"io"
	"log"
)

type STDINOperationReader struct {
	reader io.Reader
}

func NewSTDINOperationReader(reader io.Reader) *STDINOperationReader {
	return &STDINOperationReader{reader: reader}
}

func (r *STDINOperationReader) ReadOperations() ([][]domain.Operation, error) {
	var operations [][]domain.Operation
	scanner := bufio.NewScanner(r.reader)
	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Text() == "" {
			break
		}

		var currentOperations []domain.Operation
		if err := json.Unmarshal([]byte(line), &currentOperations); err != nil {
			log.Printf("error unmarshalling operations: %v", err)
			return nil, err
		}

		operations = append(operations, currentOperations)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}
