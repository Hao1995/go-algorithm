package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func searchInBT(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if node.Value == val {
		return true
	} else if val > node.Value {
		return searchInBT(node.Right, val)
	} else if val < node.Value {
		return searchInBT(node.Left, val)
	}

	return false
}

// TestBinaryTreeInsert tests BinaryTree
func TestBinaryTreeInsert(t *testing.T) {

	tests := []struct {
		target    int
		setUp     func() *Node
		checkFunc func(node *Node, val int) bool
		expRes    bool
	}{
		{
			target: 14,
			setUp: func() *Node {
				// 10 -> 1, 15
				// 1 -> nil, 2
				// 15 -> 13, 16
				node := &Node{
					Value: 10,
					Left: &Node{
						Value: 1,
						Right: &Node{
							Value: 2,
						},
					},
					Right: &Node{
						Value: 15,
						Left: &Node{
							Value: 13,
						},
						Right: &Node{
							Value: 16,
						},
					},
				}

				return node
			},
			checkFunc: searchInBT,
			expRes:    true,
		},
	}

	for _, test := range tests {
		root := test.setUp()

		root.Insert(test.target)

		actual := test.checkFunc(root, test.target)

		assert.Equal(t, test.expRes, actual)
	}
}
