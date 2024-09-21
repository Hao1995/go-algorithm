package kokoeatingbananas

import "math"

func minEatingSpeed(piles []int, h int) int {
	// O(n)
	var maxPile int
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	calTotalHours := func(eatingRate int) int {
		var ans int
		// O(n)
		for _, pile := range piles {
			ans += int(math.Ceil(float64(pile) / float64(eatingRate)))
		}
		return ans
	}

	// O(logn*n) = O(nlogn)
	var l, r int = 1, maxPile + 1
	for l < r {
		var midd int = (l + r) / 2

		totalHours := calTotalHours(midd)
		if totalHours > h {
			// invalid case, try faster speed
			l = midd + 1
		} else {
			// valid case, try slower speed
			r = midd
		}
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
