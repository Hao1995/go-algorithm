package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{-1, head}
	fastNode, slowNode := dummyNode, dummyNode

	// move fastNode to the end and slowNode to the previous node of the target node
	for fastNode.Next != nil {
		if n <= 0 {
			slowNode = slowNode.Next
		}
		fastNode = fastNode.Next
		n--
	}

	nextNode := slowNode.Next
	slowNode.Next = nextNode.Next

	return dummyNode.Next
}
