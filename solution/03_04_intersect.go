package solution

// 37%
func Intersect(K int, L int, M int, N int, P int, Q int, R int, S int) bool {
	// write your code in Go 1.4
	a := (float64(N) - float64(L)) / (float64(M) - float64(K))
	b := float64(L) - a*float64(K)

	PQ := a*float64(P) + b - float64(Q)
	RS := a*float64(R) + b - float64(S)

	if PQ == 0 || RS == 0 {
		return true
	}
	if PQ > 0 && RS < 0 {
		return true
	}
	if PQ < 0 && RS > 0 {
		return true
	}

	return false
}
