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

func (b *BinarySearchTree) Search(val int) *Node {
	if b.Root == nil {
		return nil
	}

	currNode := b.Root
	for {
		if val == currNode.Value {
			return currNode
		}

		if val < currNode.Value {
			if currNode.Left == nil {
				return nil
			}

			currNode = currNode.Left

			continue
		}

		if val > currNode.Value {
			if currNode.Right == nil {
				return nil
			}

			currNode = currNode.Right

			continue
		}
	}
}

func (b *BinarySearchTree) Delete(val int) bool {
	newBST, deleted := deleteNode(b.Root, val)

	b.Root = newBST

	return deleted
}

func deleteNode(root *Node, key int) (*Node, bool) {
	var deleted bool

	if root == nil {
		return nil, deleted
	}

	if root.Value == key {
		deleted = true

		// case: no child
		if root.Right == nil && root.Left == nil {
			return nil, deleted
		}

		// case: only has right child
		if root.Right != nil && root.Left == nil {
			return root.Right, deleted
		}

		// case: only has left child
		if root.Left != nil && root.Right == nil {
			return root.Left, deleted
		}

		// case: has two child
		minTreeNode := getMinNode(root.Right)
		leftTree := root.Left
		rightTree, deleted := deleteNode(root.Right, minTreeNode.Value)
		minTreeNode.Left = leftTree
		minTreeNode.Right = rightTree
		return minTreeNode, deleted
	}

	if root.Value > key {
		root.Left, deleted = deleteNode(root.Left, key)
	}

	if root.Value < key {
		root.Right, deleted = deleteNode(root.Right, key)
	}

	return root, deleted
}

func getMinNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return getMinNode(root.Left)
}
