# Intuition
Just sort the points by closest distance in place, then return the first k points.

# Approach
---

# Complexity
- Time complexity:
I am using golang build-in sorting algorithm: O(N*log(N))

- Space complexity:
O(1) ?

# Code
```
func kClosest(points [][]int, k int) [][]int {
    sort.Slice(points, func(i, j int) bool {
        distA := points[i][0] * points[i][0] + points[i][1] * points[i][1]
        distB := points[j][0] * points[j][0] + points[j][1] * points[j][1]
        return distA < distB
    })
    
    return points[:k]
}
```