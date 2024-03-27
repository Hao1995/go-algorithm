package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	var tmpAns []int
	var queue, nextQueue []*TreeNode = []*TreeNode{root}, []*TreeNode{}
	var node *TreeNode
	var isLeft bool = true
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			queue, nextQueue = nextQueue, []*TreeNode{}
			ans, tmpAns = append(ans, tmpAns), []int{}
			isLeft = !isLeft
		}

		node, queue = queue[0], queue[1:]
		if isLeft {
			tmpAns = append(tmpAns, node.Val)
		} else {
			tmpAns = append([]int{node.Val}, tmpAns...)
		}

		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}

		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}
	}

	if len(tmpAns) > 0 {
		ans = append(ans, tmpAns)
	}

	return ans
}
