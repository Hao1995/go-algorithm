package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func traversalFunc(bst *BinarySearchTree) []int {
	if bst == nil {
		return []int{}
	}

	arr := traversal(bst.Root, []int{})
	return arr
}

func traversal(node *Node, arr []int) []int {
	if node == nil {
		return arr
	}

	arr = append(arr, node.Value)

	if node.Left != nil {
		arr = traversal(node.Left, arr)
	}
	if node.Right != nil {
		arr = traversal(node.Right, arr)
	}

	return arr
}

// TestBinaryTreeInsert tests BinaryTree
func TestBinaryTreeInsert(t *testing.T) {

	tests := []struct {
		target        int
		setUp         func(bst *BinarySearchTree)
		traversalFunc func(bst *BinarySearchTree) []int
		expRes        []int
	}{
		{
			target: 14,
			setUp: func(bst *BinarySearchTree) {
				//      10
				//   1      15
				// n   2  13   16
				bst.Insert(10)
				bst.Insert(1)
				bst.Insert(15)
				bst.Insert(2)
				bst.Insert(13)
				bst.Insert(16)
			},
			traversalFunc: traversalFunc,
			expRes:        []int{10, 1, 2, 15, 13, 16},
		},
	}

	for _, test := range tests {
		bst := &BinarySearchTree{}

		test.setUp(bst)

		actual := test.traversalFunc(bst)

		assert.DeepEqual(t, test.expRes, actual)
	}
}
