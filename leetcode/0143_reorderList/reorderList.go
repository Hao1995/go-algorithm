package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var pt1, pt2 *ListNode = head, head
	if pt1.Next == nil {
		return
	}

	// find middle Node
	for pt2.Next != nil && pt2.Next.Next != nil {
		pt2 = pt2.Next.Next
		pt1 = pt1.Next
	}
	var middle *ListNode = pt1

	// push the second half of nodes
	var backNodeStack []*ListNode
	for pt1.Next != nil {
		backNodeStack = append(backNodeStack, pt1.Next)
		pt1 = pt1.Next
	}

	// avoid circle link
	middle.Next = nil

	// arrange new list
	var ptr *ListNode = head
	for len(backNodeStack) > 0 {
		var backNode *ListNode
		backNode, backNodeStack = backNodeStack[len(backNodeStack)-1], backNodeStack[:len(backNodeStack)-1]

		frontTmpNode := ptr.Next
		ptr.Next = backNode
		backNode.Next = frontTmpNode
		ptr = frontTmpNode
	}
}
