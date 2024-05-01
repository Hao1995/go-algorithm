package findkclosestelements

func findClosestElements(arr []int, k int, x int) []int {
	var l, r int = 0, len(arr) - k
	for l < r {
		var midd int = (l + r) / 2
		if x-arr[midd] > arr[midd+k]-x {
			l = midd + 1
		} else {
			r = midd
		}
	}
	return arr[l : l+k]
}
