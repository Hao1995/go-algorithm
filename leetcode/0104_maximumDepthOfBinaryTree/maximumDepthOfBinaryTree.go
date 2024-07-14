package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue, nextQueue []*TreeNode = []*TreeNode{root}, []*TreeNode{}
	var depth int = 1
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			queue, nextQueue = nextQueue, []*TreeNode{}
			depth++
			continue
		}

		var node *TreeNode
		node, queue = queue[0], queue[1:]
		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}
	}

	return depth
}
