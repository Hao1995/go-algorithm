package solution

import "github.com/Hao1995/algorithm/models"

// BinaryTreeVisibleNodesCounting :
// Example
// (5, (3, (20, None, None), (21, None, None)), (10, (1, None, None), None)) >> 4 (5,20,21,10)
// (8, (2, (8, None, None), (7, None, None)), (6, None, None)) >> 2 (8,8)
func BinaryTreeVisibleNodesCounting(T *models.Tree) int {

	if T == nil {
		return 0
	}
	return binaryTreeVisibleNodesCounting(T, -1<<63)
}

func binaryTreeVisibleNodesCounting(T *models.Tree, max int) int {

	var num int
	if T.Val >= max {
		num++
		max = T.Val
	}

	if T.Left != nil {
		num += binaryTreeVisibleNodesCounting(T.Left, max)
	}

	if T.Right != nil {
		num += binaryTreeVisibleNodesCounting(T.Right, max)
	}

	return num
}
