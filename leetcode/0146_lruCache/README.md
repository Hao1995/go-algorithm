# 146. LRU Cache

## Intuition
My initial thought involves utilizing `hash` and `sort`.
However, employing `sort` (specifically quick sort) introduces a time complexity of O(nlogn).
Unfortunately, the time complexity requirement for this problem is O(1) for both `get` and `put` operations.

So I found inspiration in a solution here: `https://leetcode.com/problems/lru-cache/solutions/45911/java-hashtable-double-linked-list-with-a-touch-of-pseudo-nodes`

## Approach
To address the time complexity requirement, I opted for a solution involving `hash` and a `double linked-list`.
By leveraging the hash, we can locate the node and efficiently move it to either the head or tail of the `linked-list`.

// init, head->tail
// put 1,1, head->node(1,1)->tail
// put 2,2, head->node(2,2)->node(1,1)->tail
// get 1,   head->node(2,2)->node(1,1)->tail, return 1
// put 3,3, head->node(1,1)->node(3,3)->tail
// get 2,   head->node(1,1)->node(3,3)->tail, return -1
// put 4,4, head->node(3,3)->node(4,4)->tail
// get 1,   head->node(3,3)->node(4,4)->tail, return -1
// get 3,   head->node(3,3)->node(4,4)->tail, return 3
// get 4,   head->node(3,3)->node(4,4)->tail, return 4
