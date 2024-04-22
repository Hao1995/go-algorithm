# 377. Combination Sum IV

## Intuitive
DFS... But it will exceed the time limitation without cache.

## Approach
Dynamic programming.

To avoid the duplicated result, we must add a cache array to store the number of the combination first.
target=K, cache[K] += cache[K-nums[0]] + cache[K-nums[1]] + ... + cache[N-nums[len(nums)-1]]

For example:
nums=[1,2,3], target=4
target=4, cache[4] += cache[4-1] + cache[4-2] + cache[4-3] = cache[3] + cache[2] + cache[1]
target=3, cache[3] += cache[2] + cache[1] + cache[0]
target=2, cache[2] += cache[1] + cache[0]
target=1, cache[1] += cache[0]
target=0, cache[0] = 0