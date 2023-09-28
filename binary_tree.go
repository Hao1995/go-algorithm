package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if n == nil {
		return
	}

	if val > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	} else if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(node *Node, val int) int {
	return 0
}
