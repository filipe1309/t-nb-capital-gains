package main

import (
	calculateCapitalGain "capital-gains/internal/service/calculate-capital-gain"
	"os"
)

func main() {
	if err := calculateCapitalGain.Calculate2(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
