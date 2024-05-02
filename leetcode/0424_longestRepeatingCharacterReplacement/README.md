# 424. Longest Repeating Character Replacement

## Intuition
Hash map: Count the most frequently occurring letter.
Sliding window: Start from size 1 window, keep increase the size if current set is valid, otherwise move the window to right.

## Approach
Ref: Athfan159. https://leetcode.com/problems/longest-repeating-character-replacement/solutions/2905786/needcode-s-solution-in-golang-way-with-explanation-beats-92-5
O(n).

Only a variable to store the maximum frequency.
But you may have a question: why don't we update the maximum frequency even we move the entire sliding window one step to the right?
Right, it actually possible to change. But it doesn't matter because if the frequency of the new letter doesn't exceed the previous maximum frequency, the answer would not be updated.
So if the answer would not be updated, the sliding window will keep moving to the right until the end.

// "AABABBA", k=1
// l=0, r=0, count=[A:1], mostHz=1, total=1, otherHz=1-1=0, otherHZ <= k, is valid, update ans=max(ans, r-l+1)=1
// l=0, r=1, count=[A:2], mostHz=2, total=2, otherHz = 0 <= k, is valid, update ans=2
// l=0, r=2, count=[A:2,B:1], mostHz=2, total=3, otherHz = 1 <= k, is valid, update ans=3
// l=0, r=3, count=[A:3,B:1], mostHz=3, total=4, otherHz = 1 <= k, is valid, update ans=4
// l=0, r=4, count=[A:3,B:2], mostHz=3, total=5, otherHz = 2 <= k, is invalid, increment `l`
// l=1, r=5, count=[A:2,B:3], mostHz=3, total=5, otherHz = 2 <= k, is invalid, increment `l`
// l=2, r=6, count=[A:2,B:3], mostHz=3, total=5, otherHz = 2 <= k, is invalid, increment `l`
// l=3, r=7, r is out of the boundary.

There is a easy way to understand, but it should calculate the maximum frequency every time. O(26*n): easyCharacterReplacement
// "AABABBA", k=1
// l=0, r=0, count=[A:1], mostHz=1, total=1, otherHz=1-1=0, otherHZ <= k, is valid, update ans=max(ans, r-l+1)=1
// l=0, r=1, count=[A:2], mostHz=2, total=2, otherHz = 0 <= k, is valid, update ans=2
// l=0, r=2, count=[A:2,B:1], mostHz=2, total=3, otherHz = 1 <= k, is valid, update ans=3
// l=0, r=3, count=[A:3,B:1], mostHz=3, total=4, otherHz = 1 <= k, is valid, update ans=4
// l=0, r=4, count=[A:3,B:2], mostHz=3, total=5, otherHz = 2 <= k, is invalid, add `l`, decrement count=[A:2,B:2]
// l=1, r=4, count=[A:2,B:2], mostHz=2, total=4, otherHz = 2 <= k, is invalid, add `l`, decrement count=[A:1,B:2]
// l=2, r=4, count=[A:1,B:2], mostHz=2, total=3, otherHz = 1 <= k, is valid, ans still is = 4
// l=2, r=5, count=[A:1,B:3], mostHz=3, total=4, otherHz = 1 <= k, is valid, ans still is = 4
// l=2, r=6, count=[A:2,B:3], mostHz=3, total=5, otherHz = 2 <= k, is invalid, add `l`, decrement count=[A:2,B:2]
// l=3, r=6, count=[A:2,B:2], mostHz=2, total=4, otherHz = 2 <= k, is invalid, add `l`, decrement count=[A:1,B:2]
// l=4, r=6, count=[A:1,B:2], mostHz=2, total=3, otherHz = 1 <= k, is valid, ans still is 4
// l=4, r=7, count=[A:1,B:2], mostHz=2, total=3, otherHz = 1 <= k, is valid, ans still is 4
// l=4, r=8, out of boundary.
