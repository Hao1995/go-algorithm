# 239.Sliding Window Maximum

## Intuition
n=len(nums), k=size of window
The burst solution is iterate the window numbers each move >> O(n*k)

But one of the constraints is `1 <= k <= nums.length`, the worst case would be O(n^2). (k=n)

The maximum `n` is 10^5, we can't run O(n^2) efficiently.
Can I improve the solution to O(nlogn)? The only thoughts in my mind is priority queue.
But it's too hard to considerate which `num` should be remove from the queue.

## Approach
Solution from NeetCode: https://youtu.be/DfljaUwZsOk?si=CaPiKng8Dc7Bddsl

We don't need to store all numbers within the window into the PQ, just need to store the numbers that smaller than the maximum number of the window.
Firstly, create an array to store the numbers with decreasing order, remove all numbers less than the upcoming number. (Best case: O(n), Worst case: O(n^2))

// nums=[1,3,-1,-3,5,3,2,2], k=3, q should store idx of nums
// l=0, r=0, num=nums[0]=1, append, q=[1], output=[]
// l=0, r=1, num=3, remove smaller nums, q=[3], output=[]
// l=0, r=2, num=-1, append, q=[3,-1], r-l+1 == k, add to output=[3]
// l=1, r=3, num=-3, append, q=[3,-1,-3], r-l+1 == k, add to output=[3,3]
// l=2, r=4, num=5, remove smaller nums, q=[5], r-l+1 == k, add to output=[3,3,5]
// l=3, r=5, num=3, append, q=[5,3], r-l+1 == k, add to output=[3,3,5,5]
// l=4, r=6, num=2, append, q=[5,3,2], r-l+1 == k, add to output=[3,3,5,5,5]
// l=5, r=7, num=2, append, q=[5,3,2,2], l>q[0], remove, q=[3,2,2], r-l+1 == k, add to output=[3,3,5,5,5,3]