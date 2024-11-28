package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(pre []int, in []int) *TreeNode
	dfs = func(pre []int, in []int) *TreeNode {
		if len(pre) == 0 || len(in) == 0 {
			return nil
		}

		rootNum := pre[0]

		var rootIdxOfIn int
		for idx, num := range in {
			if num == rootNum {
				rootIdxOfIn = idx
				break
			}
		}

		return &TreeNode{
			Val:   rootNum,
			Left:  dfs(pre[1:1+rootIdxOfIn], in[:rootIdxOfIn]),
			Right: dfs(pre[1+rootIdxOfIn:], in[rootIdxOfIn+1:]),
		}
	}

	return dfs(preorder, inorder)
}
