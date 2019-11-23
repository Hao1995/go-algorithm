package solution

import "testing"

func TestTwoRectAreaFixed(t *testing.T) {
	var ans int
	var x11, y11, x12, y12, x21, y21, x22, y22 int
	x11, y11, x12, y12, x21, y21, x22, y22 = 0, 0, 5, 5, 3, -2, 6, 1
	ans = 32
	if val := TwoRectAreaFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != ans {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect %v, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, ans, val)
	}

	x11, y11, x12, y12, x21, y21, x22, y22 = 10, 10, 15, 15, 13, 7, 17, 9
	ans = 33
	if val := TwoRectAreaFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != ans {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect %v, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, ans, val)
	}

	x11, y11, x12, y12, x21, y21, x22, y22 = 7, 9, 17, 19, 9, 12, 14, 16
	ans = 100
	if val := TwoRectAreaFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != ans {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect %v, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, ans, val)
	}

	x11, y11, x12, y12, x21, y21, x22, y22 = -5, -3, 0, 7, -3, 0, 7, 4
	ans = 78
	if val := TwoRectAreaFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != ans {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect %v, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, ans, val)
	}

	t.Logf("Success!")
}
