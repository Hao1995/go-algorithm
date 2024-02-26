package rotatearray

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// // V2
// func rotate(nums []int, k int)  {
//     nk := k % len(nums)
//     gap := len(nums) - nk
//     leftPart, rightPart := make([]int, nk), make([]int, gap)

//     for i, num := range nums {
//         if i < gap {
//             rightPart[i] = num
//         } else {
//             leftPart[i - gap] = num
//         }
//     }

//     for i := range nums {
//         if i < nk  {
//             nums[i] = leftPart[i]
//         } else {
//             nums[i] = rightPart[i-nk]
//         }
//     }
// }
