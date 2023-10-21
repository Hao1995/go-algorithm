package datastructure

import (
	"reflect"
	"testing"
)

func TestAVLTreeInsert(t *testing.T) {
	tests := []struct {
		desc      string
		tree      *AVLTree
		input     int
		expResult []int
	}{
		{
			desc: "with right and left children",
			tree: &AVLTree{
				Root: &AVLNode{
					Val: 10,
					Left: &AVLNode{
						Val: 7,
						Left: &AVLNode{
							Val: 2,
						},
						Right: &AVLNode{
							Val: 8,
						},
					},
					Right: &AVLNode{
						Val: 13,
						Left: &AVLNode{
							Val: 11,
						},
						Right: &AVLNode{
							Val: 18,
						},
					},
				},
			},
			input:     15,
			expResult: []int{2, 7, 8, 10, 11, 13, 15, 18},
		},
		{
			desc: "without right and left children",
			tree: &AVLTree{
				Root: &AVLNode{
					Val: 10,
				},
			},
			input:     15,
			expResult: []int{10, 15},
		},
		{
			desc: "tree is null",
			tree: &AVLTree{
				Root: nil,
			},
			input:     15,
			expResult: []int{15},
		},
	}

	for _, test := range tests {
		test.tree.Insert(test.input)
		actualRes := avlTreeInOrder(test.tree)

		if !reflect.DeepEqual(test.expResult, actualRes) {
			t.Fatalf(`Insert failed. expRes=%v, actualRes=%v`, test.expResult, actualRes)
		}
	}
}

func avlTreeInOrder(tree *AVLTree) []int {
	if tree == nil {
		return []int{}
	}

	if tree.Root == nil {
		return []int{}
	}

	return traversal(tree.Root)
}

func traversal(node *AVLNode) []int {
	ans := []int{}
	if node == nil {
		return ans
	}

	if node.Left != nil {
		ans = append(ans, traversal(node.Left)...)
	}

	ans = append(ans, node.Val)

	if node.Right != nil {
		ans = append(ans, traversal(node.Right)...)
	}

	return ans
}
