package main

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func inOrderTrvsFunc(bst *BinarySearchTree) []int {
	fmt.Printf("inOrderTrvsFunc, bt=%v\n", bst.Root)
	return traversal(bst.Root, []int{}, 0)
}

func traversal(node *Node, arr []int, count int) []int {
	count++
	if count >= 5 {
		return arr
	}

	fmt.Printf("node=%v, arr=%v\n", node, arr)
	if node == nil {
		return arr
	}

	arr = append(arr, node.Value)

	if node.Left != nil {
		arr = traversal(node.Left, arr, count)
	}
	if node.Right != nil {
		arr = traversal(node.Right, arr, count)
	}

	return arr
}

// // TestBinaryTreeInsert tests BinaryTree
// func TestBinaryTreeInsert(t *testing.T) {
// 	tests := []struct {
// 		target        int
// 		setUp         func(bst *BinarySearchTree)
// 		traversalFunc func(bst *BinarySearchTree) []int
// 		expRes        []int
// 	}{
// 		{
// 			target: 14,
// 			setUp: func(bst *BinarySearchTree) {
// 				//      10
// 				//   1      15
// 				// n   2  13   16
// 				bst.Insert(10)
// 				bst.Insert(1)
// 				bst.Insert(15)
// 				bst.Insert(2)
// 				bst.Insert(13)
// 				bst.Insert(16)
// 			},
// 			traversalFunc: inOrderTrvsFunc,
// 			expRes:        []int{10, 1, 2, 15, 13, 16},
// 		},
// 	}

// 	for _, test := range tests {
// 		bst := &BinarySearchTree{}

// 		test.setUp(bst)

// 		actual := test.traversalFunc(bst)

// 		assert.DeepEqual(t, test.expRes, actual)
// 	}
// }

// // TestBinaryTreeSearch tests Search method in BinaryTree
// func TestBinaryTreeSearch(t *testing.T) {
// 	tests := []struct {
// 		desc          string
// 		target        int
// 		setUp         func(desc string) *BinarySearchTree
// 		traversalFunc func(bst *BinarySearchTree) []int
// 		expRes        *Node
// 	}{
// 		{
// 			desc:   "success to find value",
// 			target: 2,
// 			setUp: func(desc string) *BinarySearchTree {
// 				//      10
// 				//   1      15
// 				// n   2  13   16
// 				btree := &BinarySearchTree{}
// 				btree.Root = &Node{
// 					Value: 10,
// 					Left: &Node{
// 						Value: 1,
// 						Right: &Node{
// 							Value: 2,
// 						},
// 					},
// 					Right: &Node{
// 						Value: 15,
// 						Left: &Node{
// 							Value: 13,
// 						},
// 						Right: &Node{
// 							Value: 16,
// 						},
// 					},
// 				}
// 				return btree
// 			},
// 			expRes: &Node{
// 				Value: 2,
// 			},
// 		},
// 		{
// 			desc:   "failed to find non-exist value",
// 			target: 12,
// 			setUp: func(desc string) *BinarySearchTree {
// 				//      10
// 				//   1      15
// 				// n   2  13   16
// 				btree := &BinarySearchTree{}
// 				btree.Root = &Node{
// 					Value: 10,
// 					Left: &Node{
// 						Value: 1,
// 						Right: &Node{
// 							Value: 2,
// 						},
// 					},
// 					Right: &Node{
// 						Value: 15,
// 						Left: &Node{
// 							Value: 13,
// 						},
// 						Right: &Node{
// 							Value: 16,
// 						},
// 					},
// 				}
// 				return btree
// 			},
// 			expRes: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		bt := test.setUp(test.desc)

// 		actual := bt.Search(test.target)

// 		assert.DeepEqual(t, test.expRes, actual)
// 	}
// }

// TestBinaryTreeSearch tests Search method in BinaryTree
func TestBinaryTreeDelete(t *testing.T) {
	tests := []struct {
		desc      string
		target    int
		setUp     func(desc string) *BinarySearchTree
		checkFunc func(desc string, bst *BinarySearchTree)
		expRes    bool
	}{
		{
			desc:   "success to delete the root node",
			target: 10,
			setUp: func(desc string) *BinarySearchTree {
				//      10
				//   1      15
				// n   2  13   16
				btree := &BinarySearchTree{}
				btree.Root = &Node{
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
				return btree
			},
			//       13
			//   1       15
			// n   2   n   16
			checkFunc: func(desc string, bt *BinarySearchTree) {
				list := inOrderTrvsFunc(bt)
				assert.DeepEqual(t, []int{13, 1, 2, 15, 16}, list)
			},
			expRes: true,
		},
		{
			desc:   "success to delete the smallest node",
			target: 1,
			setUp: func(desc string) *BinarySearchTree {
				//      10
				//   1      15
				// n   2  13   16
				btree := &BinarySearchTree{}
				btree.Root = &Node{
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
				return btree
			},
			checkFunc: func(desc string, bt *BinarySearchTree) {
				list := inOrderTrvsFunc(bt)
				assert.DeepEqual(t, []int{10, 2, 15, 13, 16}, list)
			},
			expRes: true,
		},
		{
			desc:   "delete an non-exist node",
			target: 3,
			setUp: func(desc string) *BinarySearchTree {
				//      10
				//   1      15
				// n   2  13   16
				btree := &BinarySearchTree{}
				btree.Root = &Node{
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
				return btree
			},
			checkFunc: func(desc string, bt *BinarySearchTree) {
				list := inOrderTrvsFunc(bt)
				assert.DeepEqual(t, []int{10, 1, 2, 15, 13, 16}, list)
			},
			expRes: false,
		},
		{
			desc:   "success to delete node without child",
			target: 2,
			setUp: func(desc string) *BinarySearchTree {
				//      10
				//   1      15
				// n   2  13   16
				btree := &BinarySearchTree{}
				btree.Root = &Node{
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
				return btree
			},
			checkFunc: func(desc string, bt *BinarySearchTree) {
				list := inOrderTrvsFunc(bt)
				assert.DeepEqual(t, []int{10, 1, 15, 13, 16}, list)
			},
			expRes: true,
		},
		{
			desc:   "success to delete node without right child",
			target: 15,
			setUp: func(desc string) *BinarySearchTree {
				//      10
				//   1      15
				// n   2  13   n
				btree := &BinarySearchTree{}
				btree.Root = &Node{
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
					},
				}
				return btree
			},
			checkFunc: func(desc string, bt *BinarySearchTree) {
				list := inOrderTrvsFunc(bt)
				assert.DeepEqual(t, []int{10, 1, 2, 13}, list)
			},
			expRes: true,
		},
	}

	for _, test := range tests {
		bt := test.setUp(test.desc)

		actual := bt.Delete(test.target)

		assert.Equal(t, test.expRes, actual)

		test.checkFunc(test.desc, bt)
	}
}
