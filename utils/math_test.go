package utils

import "testing"

func TestRound(t *testing.T) {

	var input float64
	var digit int
	var ans float64

	input = 10.1234
	digit = 2
	ans = 10.12
	if val := Round(input, digit); val != ans {
		t.Errorf("Except to get %v. But get %v", ans, val)
	}
}
