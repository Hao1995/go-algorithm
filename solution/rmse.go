package solution

import "math"

// RMSE :
// ([4, 25,  0.75, 11],[3, 21, -1.25, 13]) >> 2.5
// My test case
// ([0.5, 3, 3, 5.5],[1, 2, 3, 6]) >> 0.612372436
// ([-10,7,23,3,4],[13,7,30,1,8]) >> 10.936178492
// ([3,2,5,7,9],[1,8,5,7,9]) >> 2.828427125
func RMSE(predicted []float64, observed []float64) float64 {

	length := len(predicted)

	if length != len(observed) {
		return 0
	}

	var ans float64
	for i := 0; i < length; i++ {
		ans += math.Pow(predicted[i]-observed[i], 2)
	}

	ans /= float64(length)

	return math.Sqrt(ans)
}
