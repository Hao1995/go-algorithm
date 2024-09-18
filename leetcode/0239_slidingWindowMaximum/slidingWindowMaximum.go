package slidingwindowmaximum

func maxSlidingWindow(nums []int, k int) []int {
	var ans, q []int
	var l, r int

	for r < len(nums) {
		// remove smaller nums
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		// add to queue
		q = append(q, r)

		// remove the left nums that out of window
		if l > q[0] {
			q = q[1:]
		}

		// store the curr max num
		if r-l+1 == k {
			ans = append(ans, nums[q[0]])
			l++
		}

		r++
	}

	return ans
}
