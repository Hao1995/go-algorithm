package subtreeofanothertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var isSame func(node1, node2 *TreeNode) bool
	isSame = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		return isSame(node1.Left, node2.Left) && isSame(node1.Right, node2.Right)
	}

	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil {
			return false
		}

		if isSame(node1, node2) {
			return true
		}

		return dfs(node1.Left, node2) || dfs(node1.Right, node2)
	}

	return dfs(root, subRoot)
}
