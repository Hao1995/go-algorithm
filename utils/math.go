package utils

import "math"

func Round(input float64, digit int) (output float64) {

	times := math.Pow(10, float64(digit))
	output = float64(int64(input*times)) / times

	return
}
