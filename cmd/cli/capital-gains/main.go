package main

import (
	"capital-gains/internal/application"
	"capital-gains/internal/domain"
	"capital-gains/internal/infrastructure"
	"os"
)

func main() {
	calculator := domain.NewCalculator()
	operationReader := infrastructure.NewSTDINOperationReader(os.Stdin)
	taxesWriter := infrastructure.NewSTDOUTTaxesWriter()
	service := application.NewCalculatorService(calculator, operationReader, taxesWriter)

	if err := service.Calculate(); err != nil {
		panic(err)
	}
}
