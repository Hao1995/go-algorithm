package solution

import "sort"

// 75%
// case
// (-10,-5,3,4,-8,-4,0,0)
// (0,0,5,7,-2,5,2,9)
// (3,3,6,8,5,0,7,5)
// (1,3,5,5,2,4,6,7)
// (0,3,5,7,-3,2,3,5)
// (1,1,7,5,2,0,5,2)
// (3,5,10,8,5,0,7,3)
func TwoRectArea(K int, L int, M int, N int, P int, Q int, R int, S int) int {
	// write your code in Go 1.4
	var max int32
	max = -1 << 31
	max--

	rectA := float64(M-K) * float64(N-L)
	if rectA < 0 {
		return -1
	}
	rectB := float64(R-P) * float64(S-Q)
	if rectB < 0 {
		return -1
	}

	union := rectA + rectB
	if union < 0 {
		return -1
	}

	if !(P >= M || R <= K || S <= L || Q >= N) {
		// it's mean they have intersection
		x := []int{K, P, R, M}
		sort.Ints(x)

		y := []int{L, Q, S, N}
		sort.Ints(y)

		interArea := float64(x[2]-x[1]) * float64(y[2]-y[1])
		union -= interArea
	}

	return int(union)
}
