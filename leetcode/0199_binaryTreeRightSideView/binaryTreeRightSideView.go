package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var queue []*TreeNode
	var nextQueue []*TreeNode

	if root == nil {
		return []int{}
	}

	var ans []int
	ans = append(ans, root.Val)
	queue = append(queue, root)

	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			ans = append(ans, nextQueue[len(nextQueue)-1].Val)
			queue = nextQueue
			nextQueue = []*TreeNode{}
		}

		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = []*TreeNode{}
	}

	return ans
}
