package infrastructure

import (
	"capital-gains/internal/domain"
	"database/sql"
)

type DBOperationReader struct {
	db *sql.DB
}

func NewDBOperationReader(db *sql.DB) *DBOperationReader {
	return &DBOperationReader{db: db}
}

func (r *DBOperationReader) ReadOperations() ([][]domain.Operation, error) {
	var operations [][]domain.Operation

	// TODO: Implement the ReadOperations method for the DBOperationReader

	return operations, nil
}
