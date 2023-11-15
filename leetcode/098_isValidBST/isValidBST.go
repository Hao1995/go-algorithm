package isvalidbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return traversal(root, (1<<31)*-1-1, 1<<31)
}

func traversal(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minVal {
		return false
	}

	if root.Val >= maxVal {
		return false
	}

	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		if !traversal(root.Left, minVal, root.Val) {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		if !traversal(root.Right, root.Val, maxVal) {
			return false
		}
	}

	return true
}
