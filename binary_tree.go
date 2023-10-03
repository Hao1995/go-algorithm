package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (b *BinarySearchTree) Insert(val int) {
	if b.Root == nil {
		b.Root = &Node{Value: val}
		return
	}

	currNode := b.Root
	for {
		if val > currNode.Value {
			if currNode.Right == nil {
				currNode.Right = &Node{Value: val}
				return
			}
			currNode = currNode.Right
		} else if val < currNode.Value {
			if currNode.Left == nil {
				currNode.Left = &Node{Value: val}
				return
			}
			currNode = currNode.Left
		} else {
			// val is exist, return
			return
		}
	}
}

func (b *BinarySearchTree) Search(node *Node, val int) int {
	return 0
}
