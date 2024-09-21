package reorderlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderListV2(head *ListNode) {
	// find second half linked-list
	var s, f *ListNode = head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	// reverse the second-half
	shead := s.Next
	var prev *ListNode
	for shead != nil {
		tmpNode := shead.Next
		shead.Next = prev
		prev = shead
		shead = tmpNode

		// 4->5->6
		// tmpNode=5, 4->nil, prev=4, shead=5
		// tmpNode=6, 5->4, prev=5, shed=6
		// tmpNode=nil, 6->5->4, prev=6, shed=nil
	}
	shead = prev

	// avoid recycle
	s.Next = nil

	// reorder
	for head != nil && shead != nil {
		tmpNode := head.Next
		head.Next = shead
		head = tmpNode

		sTmpNode := shead.Next
		shead.Next = tmpNode
		shead = sTmpNode

		// 1->2->3, 5->4, head=1, shead=5
		// tmpNode=2, 1->5->2, head=2, sTmpNode=4, shead=4
		// tmpNode=3, 1->5->2->4->3, head=3, sTmpNode=nil, shead=nil
	}
}
