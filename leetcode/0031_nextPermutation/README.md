# 31. Next Permutation

## Intuition

## Approach

// [1,2,3,4]
// find final asc order, peakIdx=3
// swap first two nums >> swap(3,4) >> [1,2,4,3]
// [1,2,4,3]
// find final asc order, peakIdx=2
// swap first two nums >> swap(2,4) >> [1,3,4,2]
// sort after nums, i=3, pekaIdx=2, swap(4,2) >> [1,3,2,4]
// [1,3,2,4]
// find final asc order, peakIdx=3
// swap first two nums >> swap(2,4) >> [1,3,4,2]
// sort after nums, i=3, pekaIdx=3, break >> [1,3,4,2]
// [1,3,4,2]
// find final asc order, peakIdx=2
// swap largestNum to front >> peakIdx=2, largeIdx=3, 3 > 2 >>  peakIdx=2, largeIdx=2, 3 < 4, swap(3,4) >> [1,4,3,2]
// sort after nums, i=3, pekaIdx=2, swap(3,2) >> [1,4,2,3]
// [1,4,2,3]
// find final asc order, peakIdx=3
// swap nums, peakIdx=3, largeIdx=3 >> swap(2,3) >> [1,4,3,2]
// [1,4,3,2]