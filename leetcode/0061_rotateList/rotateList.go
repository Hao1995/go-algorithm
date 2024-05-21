package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find node
	node := head
	var lastNode *ListNode
	var length int
	for node != nil {
		length++
		lastNode = node
		node = node.Next
	}

	// rotateSteps
	var rotateSteps int = k % length
	if rotateSteps == 0 {
		return head
	}

	// parentNode & targetNode
	seekCount := length - rotateSteps
	parentNode := head
	targetNode := head.Next
	var rotatedCount int = 1
	for rotatedCount < seekCount {
		parentNode = parentNode.Next
		targetNode = targetNode.Next
		rotatedCount++
	}

	// remove parentNode & concate head
	parentNode.Next = nil
	lastNode.Next = head

	return targetNode
}
