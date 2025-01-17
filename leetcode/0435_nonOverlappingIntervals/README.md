# 435. Non Overlapping Intervals

## Intuition
Ref: https://youtu.be/nONCGxWoUfM?si=xAR3fLsgX5fUAkDm

## Approach
// 1---3
//.    3-4
//.  2-3
// 1-2
// [1,2], res=0, prevEnd=2
// [1,3], res=1, prevEnd=2
// [2,3], res=1, prevEnd=3
// [3,4], res=1, prevEnd=4

// 1----4
//.     4-5
//.  2-3
// 1---3
// [1,3], res=0, prevEnd=3
// [1,4], res=1, prevEnd=3
// [2,3], res=2, prevEnd=3
// [4,5], res=2, prevEnd=5

//     3-4
//.      4-5
//.  2-3
// 1-----4
// [1,4], res=0, prevEnd=4
// [2,3], res=1, prevEnd=3
// [3,4], res=1, prevEnd=4
// [4,5], res=1, prevEnd=5


//   2---4
//.      4-5
//.    3---5
// 1-2
// [1,2], res=0, prevEnd=2
// [2,4], res=0, prevEnd=4
// [3,5], res=1, prevEnd=4
// [4,5], res=1, prevEnd=5