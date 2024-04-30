package maximumwidthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type IdxNode struct {
	Node *TreeNode
	Idx  int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue, nextQueue := []IdxNode{IdxNode{Node: root, Idx: 0}}, []IdxNode{}
	var idxNode IdxNode
	var startIdx, endIdx int = -1, -1
	var ans int
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			queue, nextQueue = nextQueue, []IdxNode{}
			startIdx, endIdx = -1, -1
			continue
		}

		idxNode, queue = queue[0], queue[1:]
		if startIdx == -1 {
			startIdx = idxNode.Idx
		}
		endIdx = idxNode.Idx
		ans = max(ans, endIdx-startIdx+1)

		if idxNode.Node.Left != nil {
			nextQueue = append(nextQueue, IdxNode{
				Node: idxNode.Node.Left,
				Idx:  idxNode.Idx * 2,
			})
		}

		if idxNode.Node.Right != nil {
			nextQueue = append(nextQueue, IdxNode{
				Node: idxNode.Node.Right,
				Idx:  idxNode.Idx*2 + 1,
			})
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
