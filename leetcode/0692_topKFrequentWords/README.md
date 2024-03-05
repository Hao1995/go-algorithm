# 692. Top K Frequent Words

## Intuition
1. Count by hash map O(n)
2. Put to priority queue O(m),O(m) // (order by: freq desc, lexicographical asc)
3. Retrieve top k data from queue, O(k),O(k)


## Approach
I change step.2, I put all data into a `[]string`, then sort them by frequency and lexicographical.
So time complexity would be O(m) + O (mlogm).
The final time complexity is O(n) + O(m) + O (mlogm) + O(k) = O(n+mlogm+k)
    space complexity = O(m+k)