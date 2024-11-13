package main

import (
	"bufio"
	"capital-gains/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Text() == "" {
			break
		}

		var operations []domain.Operation
		if err := json.Unmarshal([]byte(line), &operations); err != nil {
			log.Fatalf("error unmarshalling operations: %v", err)
		}

		// Process operations and calculate taxes
		calculator := domain.NewCalculator()
		taxes := calculator.CalculateTaxes(operations)

		// Format taxes as JSON and output to stdout
		if taxesJSON, err := json.Marshal(taxes); err != nil {
			log.Fatalf("error marshalling taxes: %v", err)
		} else {
			fmt.Fprintln(os.Stdout, string(taxesJSON))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
}
