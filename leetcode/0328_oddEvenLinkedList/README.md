# 328. Odd Even Linked List

## Intuition
I have an idea that is similar to fast-slow pointers.
The difference is that even with the slow pointer, we also move 2 steps at once.


## Approach
### Method 1: Mine
// Input: 1 -> 2 -> 3 -> 4 -> 5
// Init: currNode -> 1 -> 2 -> 3 -> 4 -> 5
// round 1
// oddNode -> currNode >> oddNode -> 1 -> 2 -> 3 -> 4 -> 5
// currNode = currnode->Next >> currNode = 2 -> 3 -> 4 -> 5
// oddNode.Next = nil >> 1
// round 2
// evenNode -> currNode >> evenNode -> 2 -> 3 -> 4 -> 5
// currNode = currnode->Next >> currNode = 3 -> 4 -> 5
// round 3
// oddNode -> currNode >> oddNode -> 3 -> 4 -> 5
// currNode = currnode->Next >> currNode = 4 -> 5
// round 4
// evenNode -> currNode >> evenNode -> 4 -> 5
// currNode = currnode->Next >> currNode = 4 -> 5
// round 5
// oddNode -> currNode >> oddNode -> 5
// currNode = currnode->Next >> currNode = 5

### Method 2: harrison6
Same method, but his code is simpler.
https://leetcode.com/problems/odd-even-linked-list/solutions/78079/simple-o-n-time-o-1-space-java-solution