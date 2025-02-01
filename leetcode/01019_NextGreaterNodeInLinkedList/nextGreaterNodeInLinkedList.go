package nextgreaternodeinlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	type item struct {
		Idx int
		Val int
	}

	var nodeLen int
	currNode := head
	for currNode != nil {
		nodeLen++
		currNode = currNode.Next
	}

	ans := make([]int, nodeLen)
	var stack []item
	var idx int
	for head != nil {
		if len(stack) == 0 {
			stack = append(stack, item{idx, head.Val})
		} else {
			for len(stack) != 0 && head.Val > stack[len(stack)-1].Val {
				ans[stack[len(stack)-1].Idx] = head.Val
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, item{idx, head.Val})
		}
		idx++
		head = head.Next
	}

	return ans
}
