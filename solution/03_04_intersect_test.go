package solution

import "testing"

func TestIntersectFixed(t *testing.T) {
	var x11, y11, x12, y12, x21, y21, x22, y22 int
	x11, y11, x12, y12, x21, y21, x22, y22 = 0, 0, 3, 1, 2, 3, 3, 0
	if val := IntersectFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != true {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect true, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, val)
	}

	x11, y11, x12, y12, x21, y21, x22, y22 = 0, 0, 3, 1, 3, 3, 4, 0
	if val := IntersectFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != false {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect false, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, val)
	}

	x11, y11, x12, y12, x21, y21, x22, y22 = 1, 1, 2, 2, 2, 0, 3, 1
	if val := IntersectFixed(x11, y11, x12, y12, x21, y21, x22, y22); val != false {
		t.Errorf("(%v,%v), (%v,%v), (%v,%v), (%v,%v). Expect false, but get %v", x11, y11, x12, y12, x21, y21, x22, y22, val)
	}
}
