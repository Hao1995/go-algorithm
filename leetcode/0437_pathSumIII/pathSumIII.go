package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// sum:count
	preSumMap := make(map[int]int)
	preSumMap[0] = 1

	var ans int
	dfs(targetSum, root, 0, preSumMap, &ans)
	return ans
}

func dfs(targetSum int, node *TreeNode, sum int, preSumMap map[int]int, ans *int) {
	if node == nil {
		return
	}

	sum += node.Val
	diff := sum - targetSum
	count, ok := preSumMap[diff]
	if ok {
		*ans += count
	}
	preSumMap[sum]++

	if node.Left != nil {
		dfs(targetSum, node.Left, sum, preSumMap, ans)
	}

	if node.Right != nil {
		dfs(targetSum, node.Right, sum, preSumMap, ans)
	}

	preSumMap[sum]--
}
