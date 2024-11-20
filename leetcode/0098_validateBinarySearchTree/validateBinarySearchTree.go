package isvalidbst

const (
	MAX int = 1 << 31
	MIN int = -1<<31 - 1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, upper int, lower int) bool
	dfs = func(node *TreeNode, upper int, lower int) bool {
		if node == nil {
			return true
		}

		if node.Val >= upper || node.Val <= lower {
			return false
		}

		if node.Left != nil && !dfs(node.Left, min(node.Val, upper), lower) {
			return false
		}

		if node.Right != nil && !dfs(node.Right, upper, max(node.Val, lower)) {
			return false
		}

		return true
	}

	return dfs(root, MAX, MIN)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
