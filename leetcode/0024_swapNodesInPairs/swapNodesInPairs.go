package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	var p1, p2 *ListNode = head, head.Next
	dummyHead := &ListNode{Next: p1}
	prevNode := dummyHead
	for p1 != nil && p2 != nil {
		// record next node
		prevNode.Next = p2
		p2Next := p2.Next

		// swap
		p1, p2 = p2, p1

		// link next
		p1.Next = p2
		p2.Next = p2Next

		// continue next round
		prevNode = p2
		p1 = p2.Next
		if p1 == nil {
			break
		}
		p2 = p1.Next
	}
	return dummyHead.Next
}
