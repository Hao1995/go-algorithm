/*
2049. Count Nodes With the Highest Score
There is a binary tree rooted at 0 consisting of n nodes. The nodes are labeled from 0 to n - 1. You are given a 0-indexed integer array parents representing the tree, where parents[i] is the parent of node i. Since node 0 is the root, parents[0] == -1.
Each node has a score. To find the score of a node, consider if the node and the edges connected to it were removed. The tree would become one or more non-empty subtrees. The size of a subtree is the number of the nodes in it. The score of the node is the product of the sizes of all those subtrees.
Return the number of nodes that have the highest score.
*/

package counthighestscorenodes

func countHighestScoreNodes(parents []int) int {
	childrenMap := make(map[int][]int)
	productMap := make(map[int]int64, len(parents))

	// hash table: map[parent]children
	for i := 1; i < len(parents); i++ {
		parent := parents[i]
		childrenMap[parent] = append(childrenMap[parent], i)
	}

	// all nodes
	parentsLen := len(parents)

	// dfs. calculate the number and the product of all nodes
	dfs(0, parentsLen, childrenMap, productMap)

	// find the node which has maximum product
	var maxProduct int64 = -1
	var count int
	for _, product := range productMap {
		if product > maxProduct {
			maxProduct = product
			count = 1
			continue
		}

		if product == maxProduct {
			count++
		}
	}

	return count
}

func dfs(node int, parentsLen int, childrenMap map[int][]int, productMap map[int]int64) int {
	var sum int
	childSumList := make([]int, len(childrenMap[node]))
	for i, child := range childrenMap[node] {
		childSum := dfs(child, parentsLen, childrenMap, productMap)
		sum += childSum
		childSumList[i] = childSum
	}

	// `1` is current node
	sum += 1

	var product int64 = 1
	// * left * right
	for _, childSum := range childSumList {
		product *= int64(childSum)
	}
	// * up. up = all node - current sum
	// root node will has 0 up size
	if upSize := parentsLen - sum; upSize > 0 {
		product *= int64(upSize)
	}

	productMap[node] = product

	return sum
}
