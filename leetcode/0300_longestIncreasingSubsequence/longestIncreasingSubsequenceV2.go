package longestincreasingsubsequence

func lengthOfLISV2(nums []int) int {
	var tmp []int
	// for loop: O(n)
	for _, num := range nums {
		if size := len(tmp); size == 0 {
			tmp = append(tmp, num)
		} else {
			if num > tmp[size-1] {
				tmp = append(tmp, num)
			} else {
				// binary search: O(logn)
				// find the min `num`` which gte to current `num`
				var start, end int = 0, size - 1
				for start < end {
					var midd int = (start + end) / 2

					if tmp[midd] < num {
						start = midd + 1
					} else {
						end = midd
					}
				}
				tmp[start] = num
			}
		}
	}
	return len(tmp)
}
