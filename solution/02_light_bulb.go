package solution

// Test Cases :
// [3,2,1,5,4] > 3

func LightBulb(A []int) int {
	// write your code in Go 1.4
	count := 0
	for idx, val := range A {
		if val <= idx+1 {
			count++
		}
	}

	return count
}
