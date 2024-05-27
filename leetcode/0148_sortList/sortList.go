package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Get middle node
	middNode := getMiddle(head)
	left, right := head, middNode.Next
	middNode.Next = nil

	// Divide
	left = sortList(left)
	right = sortList(right)

	// Merge
	return merge(left, right)
}

func getMiddle(head *ListNode) *ListNode {
	var slow, fast *ListNode = head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}
