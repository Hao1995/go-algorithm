package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	dfs(root, targetSum, &ans, []int{}, 0)
	return ans
}

func dfs(node *TreeNode, targetSum int, ans *[][]int, tmpPath []int, sum int) {
	if node == nil {
		return
	}

	tmpPath = append(tmpPath, node.Val)
	sum += node.Val

	if node.Left != nil {
		dfs(node.Left, targetSum, ans, append([]int{}, tmpPath...), sum)
	}

	if node.Right != nil {
		dfs(node.Right, targetSum, ans, append([]int{}, tmpPath...), sum)
	}

	if node.Left == nil && node.Right == nil && sum == targetSum {
		*ans = append(*ans, tmpPath)
	}
}
