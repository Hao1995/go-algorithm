# 143. Reorder List

## Intuition
Linked List x Two Pointers

## Approach
1. Find the middle node using two pointers
2. Temporary store the second half of nodes.
3. Remove the `Next` node from the middle node.
4. Arrange new ListNode from `head` and the second half of nodes.

Time Complexity = O(n)
Space Complexity = O(n)