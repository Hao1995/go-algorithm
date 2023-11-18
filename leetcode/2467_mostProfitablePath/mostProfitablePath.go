/*
There is an undirected tree with n nodes labeled from 0 to n - 1, rooted at node 0. You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

At every node i, there is a gate. You are also given an array of even integers amount, where amount[i] represents:

the price needed to open the gate at node i, if amount[i] is negative, or,
the cash reward obtained on opening the gate at node i, otherwise.
The game goes on as follows:

Initially, Alice is at node 0 and Bob is at node bob.
At every second, Alice and Bob each move to an adjacent node. Alice moves towards some leaf node, while Bob moves towards node 0.
For every node along their path, Alice and Bob either spend money to open the gate at that node, or accept the reward. Note that:
If the gate is already open, no price will be required, nor will there be any cash reward.
If Alice and Bob reach the node simultaneously, they share the price/reward for opening the gate there. In other words, if the price to open the gate is c, then both Alice and Bob pay c / 2 each. Similarly, if the reward at the gate is c, both of them receive c / 2 each.
If Alice reaches a leaf node, she stops moving. Similarly, if Bob reaches node 0, he stops moving. Note that these events are independent of each other.
Return the maximum net income Alice can have if she travels towards the optimal leaf node.
*/

package mostprofitablepath

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	// init nodes relationship
	nodeLinkedMap := make(map[int][]int, len(amount))
	for _, edge := range edges {
		nodeLinkedMap[edge[0]] = append(nodeLinkedMap[edge[0]], edge[1])
		nodeLinkedMap[edge[1]] = append(nodeLinkedMap[edge[1]], edge[0])
	}

	// bobDFS: calculate the distance between bob and each node.
	var bobDis []int = make([]int, len(amount))
	for i := range bobDis {
		// `MaxInt32 / 2` is in order to greater than 10^5 (the max value of n)
		bobDis[i] = math.MaxInt32 / 2
	}
	bobDFS(bob, bobDis, nodeLinkedMap, 0, -1)

	// maxProfitDFS: calculate the max profie of Alice
	return maxProfitDFS(bob, bobDis, amount, nodeLinkedMap, 0, -1, 0)
}

func bobDFS(bob int, bobDis []int, nodeLinkedMap map[int][]int, currNode int, parentNode int) {
	if currNode == bob {
		bobDis[currNode] = 0
		return
	}

	var toBob int = math.MaxInt32 / 2
	for _, linkedNode := range nodeLinkedMap[currNode] {
		if linkedNode == parentNode {
			continue
		}

		bobDFS(bob, bobDis, nodeLinkedMap, linkedNode, currNode)
		// if the dist of linkedNode is unreachable, it will be the init val `MaxInt32/2`
		// but we already add 1, so if it's unreachable, keep `MaxInt32 / 2`
		toBob = min(toBob, bobDis[linkedNode]+1)
	}
	bobDis[currNode] = toBob
}

func maxProfitDFS(bob int, bobDis []int, amount []int, nodeLinkedMap map[int][]int, currNode int, parentNode int, step int) int {
	var maxSubSum int = math.MinInt32
	for _, linkedNode := range nodeLinkedMap[currNode] {
		if linkedNode == parentNode {
			continue
		}

		subSum := maxProfitDFS(bob, bobDis, amount, nodeLinkedMap, linkedNode, currNode, step+1)
		maxSubSum = max(maxSubSum, subSum)
	}

	var sum int
	if bobStep := bobDis[currNode]; step == bobStep {
		sum += amount[currNode] / 2
	} else if step < bobStep {
		sum += amount[currNode]
	}

	if maxSubSum != math.MinInt32 {
		sum += maxSubSum
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
