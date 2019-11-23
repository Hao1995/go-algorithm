package solution

import (
	"testing"

	"github.com/Hao1995/algorithm/utils"
)

// TestRMSE :
// ([4, 25,  0.75, 11],[3, 21, -1.25, 13]) >> 2.5
// My test case
// ([0.5, 3, 3, 5.5],[1, 2, 3, 6]) >> 0.612372
// ([-10,7,23,3,4],[13,7,30,1,8]) >> 10.936178
// ([3,2,5,7,9],[1,8,5,7,9]) >> 2.828427
func TestRMSE(t *testing.T) {

	const digits = 6

	var predicted, observed []float64
	var res float64

	predicted = []float64{4, 25, 0.75, 11}
	observed = []float64{3, 21, -1.25, 13}
	res = 2.5
	if val := utils.Round(RMSE(predicted, observed), digits); val != res {
		t.Errorf("(%v,%v). Expect %v, but get %v", predicted, observed, res, val)
	}

	predicted = []float64{0.5, 3, 3, 5.5}
	observed = []float64{1, 2, 3, 6}
	res = 0.612372
	if val := utils.Round(RMSE(predicted, observed), digits); val != res {
		t.Errorf("(%v,%v). Expect %v, but get %v", predicted, observed, res, val)
	}

	predicted = []float64{-10, 7, 23, 3, 4}
	observed = []float64{13, 7, 30, 1, 8}
	res = 10.936178
	if val := utils.Round(RMSE(predicted, observed), digits); val != res {
		t.Errorf("(%v,%v). Expect %v, but get %v", predicted, observed, res, val)
	}

	predicted = []float64{3, 2, 5, 7, 9}
	observed = []float64{1, 8, 5, 7, 9}
	res = 2.828427
	if val := utils.Round(RMSE(predicted, observed), digits); val != res {
		t.Errorf("(%v,%v). Expect %v, but get %v", predicted, observed, res, val)
	}
}
