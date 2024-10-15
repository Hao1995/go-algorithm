package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		var leftLen, rightLen int
		if node.Left != nil {
			leftLen = dfs(node.Left) + 1
		}

		if node.Right != nil {
			rightLen = dfs(node.Right) + 1
		}

		ans = max(ans, leftLen+rightLen)

		return max(leftLen, rightLen)
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
