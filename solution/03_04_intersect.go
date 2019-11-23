package solution

// 37%
func Intersect(K int, L int, M int, N int, P int, Q int, R int, S int) bool {
	// write your code in Go 1.4
	// ax + b = y
	// a(slope) = (y2-y1) / (x2-x1)
	// b = y - ax
	a := (float64(N) - float64(L)) / (float64(M) - float64(K))
	b := float64(L) - a*float64(K)

	PQ := a*float64(P) + b - float64(Q)
	RS := a*float64(R) + b - float64(S)

	if PQ == 0 || RS == 0 {
		return true
	}
	if PQ > 0 && RS < 0 {
		return true
	}
	if PQ < 0 && RS > 0 {
		return true
	}

	return false
}

func IntersectFixed(K int, L int, M int, N int, P int, Q int, R int, S int) bool {

	// L1
	// (x11,y11) = (K,L)
	// (x12,y12) = (M,N)
	// L2
	// (x21,y21) = (P,Q)
	// (x22,y22) = (R,S)

	// ax + b = y
	// a(slope) = (y2-y1) / (x2-x1)
	// b = y - ax

	// L1
	var a1, b1 float64
	if M == K {
		a1 = -1 << 63
		b1 = 0
	} else {
		a1 = (float64(N) - float64(L)) / (float64(M) - float64(K))
		b1 = float64(L) - a1*float64(K)
	}

	// L2
	var a2, b2 float64
	if R == P {
		a2 = -1 << 63
		b2 = 0
	} else {
		a2 = (float64(S) - float64(Q)) / (float64(R) - float64(P))
		b2 = float64(Q) - a2*float64(P)
	}

	if a1 == a2 {
		// parallel line
		return false
	}

	interX := (b2 - b1) / (a1 - a2)
	if ((float64(M) <= interX && interX <= float64(K)) || (float64(K) <= interX && interX <= float64(M))) && ((float64(P) <= interX && interX <= float64(R)) || (float64(R) <= interX && interX <= float64(P))) {
		// intersect point is on the L1
		return true
	}

	return false
}
