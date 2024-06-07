package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue, nextQueue []*TreeNode = []*TreeNode{root}, []*TreeNode{}
	var ans [][]int
	var subAns []int
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			queue, nextQueue = nextQueue, []*TreeNode{}
			ans = append(ans, subAns)
			subAns = []int{}
			continue
		}

		var node *TreeNode
		node, queue = queue[0], queue[1:]
		subAns = append(subAns, node.Val)
		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}
	}

	if len(subAns) != 0 {
		ans = append(ans, subAns)
	}

	return ans
}
