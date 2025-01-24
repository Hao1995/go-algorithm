package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	mapping := make(map[int]*Node)

	var dfs func(curr *Node) *Node
	dfs = func(curr *Node) *Node {
		if newNode, found := mapping[curr.Val]; found {
			return newNode
		}

		newNode := &Node{
			Val:       curr.Val,
			Neighbors: []*Node{},
		}
		mapping[curr.Val] = newNode

		for _, nextNode := range curr.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(nextNode))
		}

		return newNode
	}

	return dfs(node)
}
