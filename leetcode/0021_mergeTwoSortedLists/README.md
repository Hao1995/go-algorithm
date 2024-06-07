# 21. Merge Two Sorted Lists

## Intuition
Because both of linked lists are sorted, we can iterate them to find the smaller node.

## Approach
// l1=1->2->4, l2=1->3->4
// create a dummy and a tail node
// for loop to compare two lists, then connect tail to the smaller one
// finally, return dummy.next
