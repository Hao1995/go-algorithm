# 560. Subarray Sum Equals K

# Intuition
DFS or Sliding Window.
The time complexity of the DFS must be O(n^2).
The sliding window can not solve the problem because the numbers maybe negative, so I don't know whether should I extend the window or shrink it.

# Approach
Ref: NeetCode. https://youtu.be/fFVZt-6sgyo?si=-NMn4m9uine8AVq3
Time Complexity: O(n)

Sum up each number and count it into a prefixSumMap.
If the prefixSum exists in the map, it's a valid contiguous subarray, so add the count into the answer.
Be careful, there is an edge case. The prefixSum:0 must count 1 at beginning because we calculate the answer based on the previous existing count.

// nums=[1,-1,1,1,1,1], k=3, prefixSumMap=[0:1]
// sum=0+1=1,       diff=1-3=-2,    prefixSumMap=[0:1, 1:1]
// sum=1+(-1)=0,    diff=0-3=-3,    prefixSumMap=[0:2, 1:1]
// sum=0+1=1,       diff=1-3=-2,    prefixSumMap=[0:2, 1:2]
// sum=1+1=2,       diff=2-3=-1,    prefixSumMap=[0:2, 1:2, 2:1]
// sum=2+1=3,       diff=3-3=0,     prefixSumMap=[0:2, 1:2, 2:1],   answer+=map[0]=2
// sum=3+1=4,       diff=4-3=1,     prefixSumMap=[0:2, 1:2, 2:1],   answer+=map[1]=4