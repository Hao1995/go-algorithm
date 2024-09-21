# 143. Reorder List

## Intuition
Linked List x Two Pointers

## Approach

EX1: Store nodes of the second-half into an array.
1. Find the middle node using two pointers
2. Temporary store the second half of nodes.
3. Remove the `Next` node from the middle node.
4. Arrange new ListNode from `head` and the second half of nodes.

Time Complexity = O(n)
Space Complexity = O(n)

EX2: Directly reverse the second-half linked list.
1. Find the middle node using slow and fast pointers
2. Reverse the second-half linked list.
3. Remove the final next element of the first-half linked list.
4. Reorder the linked-list.