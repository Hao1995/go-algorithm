# 435. Non Overlapping Intervals

## Intuition
Sort in ascending order according to the `end` of each interval.
Next, iterate through the intervals to calculate the overlap. If there is an overlap, increment by 1, and finally get the ans.

# Approach
// [1,100],[11,22],[1,11],[2,12]
// sort, [1,11],[2,12],[11,22],[1,100]
// [1,11], end=11
// [2,12], overlap, ans=1
// [11,22], end=22
// [1,100], overlap, ans=2