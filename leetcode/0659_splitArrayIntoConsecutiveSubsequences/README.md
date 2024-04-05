# 659. Split Array into Consecutive Subsequences

## Intuition
ref from: https://youtu.be/-Egvieoi3Sc?si=XYCUH7_jPezBMrCS

## Approach
// nums=[1,2,3,3,4,5]
// countMap={1:1, 2:1, 3:2, 4:1, 5:1}
// endMap={}
// 1, endMap={3:1}, countMap={1:0, 2:0, 3:1, 4:1, 5:1}
// 2, not exist in countMap, next
// 3, endMap={3:1, 5:1}, countMap={1:0, 2:0, 3:0, 4:0, 5:0}

// nums=[1,2,3,3,4,4,5,5]
// countMap={1:1, 2:1, 3:2, 4:2, 5:2}
// endMap={}
// 1, endMap={3:1}, countMap={1:0, 2:0, 3:1, 4:2, 5:2}
// 2, not exist in countMap, next
// 3, endMap={3:1, 5:1}, countMap={1:0, 2:0, 3:0, 4:1, 5:1}
// 3, not exist in countMap, next
// 4, endMap={3:0, 4:1, 5:1}, countMap={1:0, 2:0, 3:0, 4:0, 5:1}
// 5, endMap={3:0, 4:0, 5:2}, countMap={1:0, 2:0, 3:0, 4:0, 5:0}

// nums=[1,2,3,4,4,5]
// countMap={1:1, 2:1, 3:1, 4:2, 5:1}
// endMap={}
// 1, endMap={3:1}, countMap={1:0, 2:0, 3:0, 4:2, 5:1}
// 2, not exist in countMap, next
// 3, not exist in countMap, next
// 4, no enough length to be a subsequence.