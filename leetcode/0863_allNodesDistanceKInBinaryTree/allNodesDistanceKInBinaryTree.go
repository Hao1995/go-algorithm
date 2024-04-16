package allnodesdistancekinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	// convert to bidirectional-graph. (note: all node values are unique)
	graph := make(map[*TreeNode][]*TreeNode)

	var node *TreeNode
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]

		if node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)

			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)

			queue = append(queue, node.Right)
		}
	}

	// BFS to find the k step nodes
	queue = graph[target]
	var nextQueue []*TreeNode

	visited := make(map[*TreeNode]bool)
	visited[target] = true

	var steps int = 1

	var ans []int
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 && len(nextQueue) > 0 {
			queue = nextQueue
			nextQueue = []*TreeNode{}
			steps++
		}

		node, queue = queue[0], queue[1:]

		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = true

		if steps == k {
			ans = append(ans, node.Val)
			continue
		}

		nearByNodes, ok := graph[node]
		if !ok {
			continue
		}

		nextQueue = append(nextQueue, nearByNodes...)
	}

	return ans
}
