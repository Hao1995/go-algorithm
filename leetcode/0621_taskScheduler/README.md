# 621. Task Scheduler

## Intuition
Priority Queue.
If `n` is greater than 0, we can't execute the same task at the next time.
So, in order to complete all the tasks with the minimum execution time, we need to execute the task with the maximum number.

I have the above intuition, but I can't implement it by myself.
I found solutions from:
1. https://www.youtube.com/live/3DZE7cfgYyg?si=M_IpBHgL4QkHi88x
2. https://leetcode.com/problems/task-scheduler/solutions/3612057/golang-optimal-approach-fully-explained-easy-solution


## Approach
// [a,a,a,b,b,b], len(pq)=2
// n=0 -> [a],[b],[a],[b],[a],[b] -> 6 (1+1+1+1+1)
// n=1 -> [a,b],[a,b],[a,b] -> 6 (2+2+2)
// n=2 -> [a,b,*],[a,b,*],[a,b] -> 8 (3+3+2)
// n=3 -> [a,b,*,*],[a,b,*,*],[a,b] -> 10 (4+4+2)