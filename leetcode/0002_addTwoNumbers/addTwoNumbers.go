package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remainder, quotient int
	var dummyHead *ListNode = &ListNode{}
	var currNode *ListNode = dummyHead
	for l1 != nil || l2 != nil || quotient > 0 {
		var sum int = quotient
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		remainder = sum % 10
		quotient = sum / 10
		currNode.Next = &ListNode{Val: remainder}
		currNode = currNode.Next
	}
	return dummyHead.Next
}
