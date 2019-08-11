package solution

import "math"

// 70%
// case
// [5,7,0]
// [1,2,3,2,1,1,1,1,1,1]
// [1,2,4,2,1,1,1,1,1,1]
// [-10,5,3,1,1,2]
// [0]
// [-3,-5,-7,-2,-3]
func BestTrip(A []int) int {
	// write your code in Go 1.4

	max := -1 << 63
	for i := 0; i < len(A); i++ {
		city1 := A[i]
		for j := i; j < len(A); j++ {
			city2 := A[j]
			if val := city1 + city2 + int(math.Abs(float64(j-i))); val > max {
				max = val
			}
		}
	}

	return max
}
