package solution

import (
	"sort"

	"github.com/Hao1995/algorithm/models"
)

// 1. Count Visible Node From Binary Tree.
// 　　The visible node is greater or equal than root value.

type SeventeenLife struct {
}

func (c *SeventeenLife) CountVisibleNodes(T *models.Tree) int {
	if T == nil {
		return -1
	}

	return c.countVisible(T, T.Val)
}

func (c *SeventeenLife) countVisible(T *models.Tree, N int) int {

	if T == nil {
		return -1
	}

	count := 0
	if T.Val >= N {
		count++
	}

	if T.Left != nil {
		if tmp := c.countVisible(T.Left, N); tmp != -1 {
			count += tmp
		}
	}

	if T.Right != nil {
		if tmp := c.countVisible(T.Right, N); tmp != -1 {
			count += tmp
		}
	}

	return count
}

// 2. Sort the number
// 　　If greater than 100,000,000. Return -1

func (c *SeventeenLife) Solution(N int) int {
	// write your code in Go 1.4

	numMap := make(map[int]int)
	count := 0
	for N != 0 {
		numMap[N%10]++
		count++
		N /= 10
	}

	keys := []int{}
	for key := range numMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	// Check whether greather than 100,000,000
	if count >= 9 {
		if val, ok := numMap[0]; ok && val < 8 {
			return -1
		}
	}

	RN := 0
	for _, key := range keys {
		for {
			if val, ok := numMap[key]; ok && val > 0 {
				RN = RN*10 + key
				numMap[key]--
				continue
			}
			break
		}
	}

	return RN
}

// pascal
// monkey cross river
