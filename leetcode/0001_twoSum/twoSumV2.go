package twosum

import "sort"

func twoSumV2(nums []int, target int) []int {
	// Change data structure
	type item struct {
		Num int
		Idx int
	}

	data := make([]item, 0, len(nums))
	for idx, num := range nums {
		data = append(data, item{num, idx})
	}

	// sort nums >> O(nlogn)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Num < data[j].Num
	})
	// fmt.Printf("data=%+v\n", data)

	// two pointer >> O(n)
	var left, right int = 0, len(data) - 1
	for left <= right {
		sum := data[left].Num + data[right].Num
		if sum == target {
			return []int{data[left].Idx, data[right].Idx}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
