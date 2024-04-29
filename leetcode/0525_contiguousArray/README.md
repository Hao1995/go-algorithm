# 525. Contiguous Array

## Intuition
Hash map, but I don't know how to calculate the maximum length of the contiguous array when the sub-array is not start at the beginning.

## Approach
Ref: https://youtu.be/agB1LyObUNE?si=5tATNj9ahkDN98_8
This solution is a hash map too.
But he store different data into the hash map. It's data structure is `diff:index`.
He calculates the difference of one and zero, then store the current index into the hash map.
If the difference is exist within the hash map, the index won't be update, because the question is aim to get the `maximum length`.

For example:
// idx       0, 1, 2, 3, 4, 5, 6, 7
// diff     -1,-2,-1,-2,-3,-4,-3,-2
// nums    [ 0, 0, 1, 0, 0, 0, 1, 1]

I got the same difference `-2` at index 1 and 3.
That mean the numbers between 1 and 3 can be a valid contiguous array.

Similarity, the index 1 and 7, with a difference of `-2` can also form a contiguous array.

Let's look how it works.

// nums=[ 0, 0, 1, 0, 0, 0, 1, 1], map[diff:idx]
// idx=0, num=0, zero=1, one=0, map=[-1:0], ans=0
// idx=1, num=0, zero=2, one=0, map=[-2:1,-1:0], ans=0
// idx=2, num=1, zero=2, one=1, map=[-2:1,-1:0], tmpAns=2-0=2, ans=2
// idx=3, num=0, zero=3, one=1, map=[-2:1,-1:0], tmpAns=3-1=2, ans=2
// idx=4, num=0, zero=4, one=1, map=[-3:4,-2:1,-1:0], tmpAns=4-4=0, ans=2
// idx=5, num=0, zero=5, one=1, map=[-4:5,-3:4,-2:1,-1:0], tmpAns=5-5=0, ans=2
// idx=6, num=1, zero=5, one=2, map=[-4:5,-3:4,-2:1,-1:0], tmpAns=6-4=2, ans=2
// idx=7, num=1, zero=5, one=3, map=[-4:5,-3:4,-2:1,-1:0], tmpAns=7-1=6, ans=6