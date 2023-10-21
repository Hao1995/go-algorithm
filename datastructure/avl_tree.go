package datastructure

type AVLTree struct {
	Root *AVLNode
}

type AVLNode struct {
	Val   int
	Left  *AVLNode
	Right *AVLNode
}

func (t *AVLTree) Insert(val int) {
	if t.Root == nil {
		t.Root = &AVLNode{
			Val: val,
		}
		return
	}

	// insert
	var currNode *AVLNode = t.Root
	for {
		if currNode.Val == val {
			rightNode := currNode.Right
			currNode.Right = &AVLNode{
				Val:   val,
				Right: rightNode,
			}
			return
		} else if val < currNode.Val {
			if currNode.Left == nil {
				currNode.Left = &AVLNode{
					Val: val,
				}
				return
			}
			currNode = currNode.Left
		} else if val > currNode.Val {
			if currNode.Right == nil {
				currNode.Right = &AVLNode{
					Val: val,
				}
				return
			}
			currNode = currNode.Right
		}
	}

	// balance
}

func (t *AVLTree) Search(val int) *AVLNode {
	return nil
}
