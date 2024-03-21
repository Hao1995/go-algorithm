package maxheap

import "fmt"

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) insert(item int) {
	h.data = append(h.data, item)
	// swap process
	currIdx := len(h.data) - 1
	for {
		var parentIdx int = (currIdx - 1) / 2
		if h.data[currIdx] > h.data[parentIdx] {
			h.data[currIdx], h.data[parentIdx] = h.data[parentIdx], h.data[currIdx]
			currIdx = parentIdx
		} else {
			break
		}
	}
}

func (h *MaxHeap) extract() int {
	if len(h.data) == 0 {
		return 0
	}

	if len(h.data) == 1 {
		ans := h.data[0]
		h.data = []int{}
		return ans
	}

	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]

	var ans int = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	var currIdx int = 0
	for currIdx < len(h.data) {
		leftIdx, rightIdx := currIdx*2+1, currIdx*2+2

		if leftIdx < len(h.data) {
			var tmpIdx int = currIdx
			if h.data[leftIdx] > h.data[currIdx] {
				tmpIdx = leftIdx
			}

			if rightIdx < len(h.data) {
				if h.data[rightIdx] > h.data[tmpIdx] {
					tmpIdx = rightIdx
				}
			}

			if currIdx == tmpIdx {
				break
			}

			h.data[tmpIdx], h.data[currIdx] = h.data[currIdx], h.data[tmpIdx]
			currIdx = tmpIdx
			continue
		}
		break
	}

	return ans
}

func (h *MaxHeap) print() {
	var nodeNum int = 1
	var count int = 0
	for i := 0; i < len(h.data); i++ {
		fmt.Printf("%v", h.data[i])
		count++
		if count == nodeNum {
			fmt.Println()
			nodeNum *= 2
			count = 0
		} else {
			fmt.Printf(", ")
		}
	}
	if count != 0 {
		fmt.Println()
	}
	fmt.Println()
}
