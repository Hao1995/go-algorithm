package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return recoverTree(0, 0, len(inorder)-1, preorder, inorder)
}

func recoverTree(preStart int, inStart, inEnd int, preorder []int, inorder []int) *TreeNode {
	if preStart >= len(preorder) || inStart > inEnd {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: preorder[preStart],
	}
	var inIdx int = 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			inIdx = i
			break
		}
	}

	root.Left = recoverTree(
		preStart+1,
		inStart,
		inIdx-1,
		preorder,
		inorder,
	)
	root.Right = recoverTree(
		// `(inIdx - inStart)` means the size of left side from inorder. EX: 1(3) - 0(9) + 1 = 2
		// `+ 1` means the right side is starting after the left side
		preStart+(inIdx-inStart)+1,
		inIdx+1,
		inEnd,
		preorder,
		inorder,
	)
	return root
}
