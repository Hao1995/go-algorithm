package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// find p, q ancestors
	var pIsFound, qIsFound bool
	var pAncestors, qAncestors []*TreeNode
	dfs(root, p, q, pIsFound, qIsFound, &pAncestors, &qAncestors)

	// find lowest common ancestors from p and q
	var lac *TreeNode
	minSize := min(len(pAncestors), len(qAncestors))
	for i := 0; i < minSize; i++ {
		if pAncestors[i].Val == qAncestors[i].Val {
			lac = pAncestors[i]
		} else {
			break
		}
	}
	return lac
}

func dfs(node, p, q *TreeNode, pIsFound, qIsFound bool, pAncestors, qAncestors *[]*TreeNode) (bool, bool) {
	if node == nil {
		return pIsFound, qIsFound
	}

	if pIsFound && qIsFound {
		return pIsFound, qIsFound
	}

	if !pIsFound {
		*pAncestors = append(*pAncestors, node)
		if node.Val == p.Val {
			pIsFound = true
		}
	}

	if !qIsFound {
		*qAncestors = append(*qAncestors, node)
		if node.Val == q.Val {
			qIsFound = true
		}
	}

	if node.Left != nil {
		pIsFound, qIsFound = dfs(node.Left, p, q, pIsFound, qIsFound, pAncestors, qAncestors)
	}

	if node.Right != nil {
		pIsFound, qIsFound = dfs(node.Right, p, q, pIsFound, qIsFound, pAncestors, qAncestors)
	}

	if !pIsFound {
		old := *pAncestors
		*pAncestors = old[:len(old)-1]
	}

	if !qIsFound {
		old := *qAncestors
		*qAncestors = old[:len(old)-1]
	}

	return pIsFound, qIsFound
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
