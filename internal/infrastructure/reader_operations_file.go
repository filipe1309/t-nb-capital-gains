package infrastructure

import (
	"bufio"
	"capital-gains/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
)

type FileOperationReader struct {
	reader io.Reader
}

func NewFileOperationReader(inputPath string) (*FileOperationReader, error) {
	input, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	return &FileOperationReader{reader: input}, nil
}

func (r *FileOperationReader) ReadOperations() ([][]domain.Operation, error) {
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

func (r *FileOperationReader) Close() {
	if closer, ok := r.reader.(io.Closer); ok {
		closer.Close()
	}
}
