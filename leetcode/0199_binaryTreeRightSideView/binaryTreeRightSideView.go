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
	if root == nil {
		return []int{}
	}

	var ans []int
	q, nq := []*TreeNode{root}, []*TreeNode{}
	for len(q) > 0 || len(nq) > 0 {
		if len(q) == 0 {
			q, nq = nq, []*TreeNode{}
			continue
		}

		var node *TreeNode
		node, q = q[0], q[1:]
		if len(q) == 0 {
			ans = append(ans, node.Val)
		}

		if node.Left != nil {
			nq = append(nq, node.Left)
		}

		if node.Right != nil {
			nq = append(nq, node.Right)
		}
	}

	return ans
}
