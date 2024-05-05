package randompickwithweight

import "math/rand"

type Solution struct {
	accArr []int
	sum    int
}

func Constructor(w []int) Solution {
	accArr := make([]int, len(w))
	var sum int
	for i := 0; i < len(w); i++ {
		sum += w[i]
		accArr[i] = sum
	}

	return Solution{
		accArr: accArr,
		sum:    sum,
	}
}

func (this *Solution) PickIndex() int {
	randNum := rand.Intn(this.sum)
	var start, end int = 0, len(this.accArr)
	for start < end {
		var midd int = (start + end) / 2
		if this.accArr[midd] <= randNum {
			start = midd + 1
		} else {
			end = midd
		}
	}
	return start
}
