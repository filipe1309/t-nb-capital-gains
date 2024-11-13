package main

import (
	calculateCapitalGain "capital-gains/internal/service/calculate-capital-gain"
	"os"
)

func main() {
	calculator := calculateCapitalGain.NewService()
	if err := calculator.Calculate(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
