package solution

import (
	"fmt"
	"math"
)

func power11(N int) int {
	str := fmt.Sprintf("%f", math.Pow(float64(11), float64(N)))
	count := 0
	for _, c := range str {
		if string(c) == "1" {
			count++
		}
	}
	return count
}

func Power11New(N int) int {

	lastArr := []int{}
	arr := []int{}
	for i := 0; i <= N; i++ {
		if i == 0 {
			arr = []int{1}
			lastArr = arr
			continue
		}
		for j := 0; j <= i; j++ {
			if j == 0 {
				arr = []int{1}
			} else if j == i {
				arr = append(arr, 1)
				lastArr = arr
				arr = []int{}
			} else {
				arr = append(arr, lastArr[j-1]+lastArr[j])
			}
		}
	}

	count := 0
	overflow := 0
	for i := len(lastArr) - 1; i > -1; i-- {
		lastArr[i] += overflow
		overflow = lastArr[i] / 10
		lastArr[i] -= overflow * 10
		if lastArr[i] == 1 {
			count++
		}
	}

	for overflow > 0 {
		num := overflow
		overflow = num / 10
		if num > 9 {
			num = overflow*10 - num
		}
		if num == 1 {
			count++
		}
	}

	return count
}
