package pseudopalindromicpaths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return dfs(root, make(map[int]int))
}

func dfs(root *TreeNode, existMap map[int]int) int {
	if root == nil {
		return 0
	}

	existMap[root.Val]++

	var palindromeNum int
	if root.Left != nil {
		palindromeNum += dfs(root.Left, existMap)
	}

	if root.Right != nil {
		palindromeNum += dfs(root.Right, existMap)
	}

	// only check leaf nodes
	if root.Left == nil && root.Right == nil && checkPalindrome(existMap) {
		existMap[root.Val]--
		return 1
	}

	existMap[root.Val]--
	return palindromeNum
}

func checkPalindrome(nodeMap map[int]int) bool {
	var oddNum int
	for _, v := range nodeMap {
		if oddNum < 1 && v%2 == 1 {
			oddNum++
			continue
		}

		if v%2 == 0 {
			continue
		}

		return false
	}

	return true
}
