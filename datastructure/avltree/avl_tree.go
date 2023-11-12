package datastructure

type AVLTree struct {
	Root *AVLNode
}

type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

func (t *AVLTree) Insert(val int) {
	if t.Root == nil {
		t.Root = &AVLNode{
			Val:    val,
			Height: -1,
		}
		return
	}

	var currNode *AVLNode = t.Root
	for {
		if currNode.Val == val {
			rightNode := currNode.Right
			currNode.Right = &AVLNode{
				Val:   val,
				Right: rightNode,
			}
			currNode.reBalanceTree()
			return
		} else if val < currNode.Val {
			if currNode.Left == nil {
				currNode.Left = &AVLNode{
					Val: val,
				}
				currNode.reBalanceTree()
				return
			}
			currNode = currNode.Left
		} else if val > currNode.Val {
			if currNode.Right == nil {
				currNode.Right = &AVLNode{
					Val: val,
				}
				currNode.reBalanceTree()
				return
			}
			currNode = currNode.Right
		}
	}
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *AVLNode) reBalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.reCalHeight()

	b := n.Left.getHeight() - n.Right.getHeight()
	if b == -2 {
		// case: right heavy
		if n.Right.Left.getHeight() > n.Right.Right.getHeight() {
			n.Right = n.Right.rotateRight()
		}
		return n.rotateLeft()
	} else if b == 2 {
		// case: left heavy
		if n.Left.Right.getHeight() > n.Left.Left.getHeight() {
			n.Left = n.Left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *AVLNode) reCalHeight() {
	var max int
	if n.Left != nil {
		max = n.Height
	}
	if n.Right != nil {
		if n.Right.Height > max {
			max = n.Right.Height
		}
	}
	n.Height = 1 + max
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n

	n.reCalHeight()
	newRoot.reCalHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n

	n.reCalHeight()
	newRoot.reCalHeight()
	return newRoot
}

func (t *AVLTree) Search(val int) *AVLNode {
	return nil
}
