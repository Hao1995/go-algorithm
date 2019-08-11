package solution

import (
	"math"
)

func NumSquaredOperations(A, B int) int {

	max := 0
	// cacheCount := make(map[int]int)
	for i := A; i <= B; i++ {
		num := math.Sqrt(float64(i))
		if num == float64(int(num)) {
			// It's squared root number
			// if val, ok := cacheCount[]
			count := 1
			for {
				tmp := math.Sqrt(num)
				if tmp == float64(int(tmp)) {
					num = tmp
					count++
				} else {
					// fmt.Printf("Num=%v, Get=%v\n", i, count)
					if count > max {
						max = count
					}
					break
				}
			}
		}
	}

	return max
}

func NumSquaredOperationsNew(A, B int) int {

	max := 0
	cacheCount := make(map[int]int)
	for i := A; i <= B; i++ {
		num := math.Sqrt(float64(i))
		if num == float64(int(num)) {
			// It's squared root number
			count := 1
			for {
				if val, ok := cacheCount[int(num)]; ok {
					count += val
				} else if tmp := math.Sqrt(num); tmp == float64(int(tmp)) {
					num = tmp
					count++
					continue
				}
				cacheCount[i] = count
				// fmt.Printf("Num=%v, Get=%v\n", i, count)
				if count > max {
					max = count
				}
				break
			}
		}
	}

	return max
}
