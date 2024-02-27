package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var odd, even *ListNode = head, head.Next
	var evenHead *ListNode = even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead

	return head
}

// // Method 1: Mine
// func oddEvenList(head *ListNode) *ListNode {
// 	// use two pointer to extract odd and even indices nodes.
// 	var oddNodes, evenNodes *ListNode = &ListNode{Next: nil}, &ListNode{Next: nil}
// 	var oddCurrNode, evenCurrNode *ListNode = oddNodes, evenNodes
// 	var skip int = 0
// 	var currNode *ListNode = head
// 	for currNode != nil {
// 		if skip > 0 {
// 			evenCurrNode.Next = currNode
// 			evenCurrNode = evenCurrNode.Next
// 			currNode = currNode.Next
// 			evenCurrNode.Next = nil
// 			skip--
// 			continue
// 		}
// 		oddCurrNode.Next = currNode
// 		oddCurrNode = oddCurrNode.Next
// 		currNode = currNode.Next
// 		oddCurrNode.Next = nil
// 		skip = 1
// 	}

// 	if evenNodes.Next != nil {
// 		oddCurrNode.Next = evenNodes.Next
// 	}

// 	return oddNodes.Next
// }

// func printNode(node *ListNode) {
// 	var currNode *ListNode = node
// 	for currNode != nil {
// 		fmt.Printf("node=%v, ", currNode.Val)
// 		currNode = currNode.Next
// 	}
// 	fmt.Println()
// }
