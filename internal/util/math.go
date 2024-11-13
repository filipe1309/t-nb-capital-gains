package util

import "math"

func RoundTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
