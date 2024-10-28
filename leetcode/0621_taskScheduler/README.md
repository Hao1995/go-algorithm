# 621. Task Scheduler

## Intuition
Priority Queue.

## Approach
// [a,a,a,b,b,b], len(pq)=2
// n=0 -> [a],[b],[a],[b],[a],[b] -> 6 (1+1+1+1+1)
// n=1 -> [a,b],[a,b],[a,b] -> 6 (2+2+2)
// n=2 -> [a,b,*],[a,b,*],[a,b] -> 8 (3+3+2)
// n=3 -> [a,b,*,*],[a,b,*,*],[a,b] -> 10 (4+4+2)

// Input length = m
// Interval = n
// PQ size = k, k <= m
// O(k*logk) >> even we have a O(nlogk) in the first nested loop, but we finally iterate all keys in PQ, so the time complexity would be O(klogk)