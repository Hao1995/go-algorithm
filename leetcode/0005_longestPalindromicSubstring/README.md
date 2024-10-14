# 5. Longest Palindromic Substring

## Intuition
For loop each character and check the longest palindromic from that index.

## Approach

// "babad"
// midd=0,      b, valid, ans=b
//              out of range
// midd=0.5,    ba, invalid
// midd=1,      a, valid, ans=a
//              bab, valid, ans=bab
//              out of range
// midd=1.5,    ab, invalid
// midd=2,      aba, valid
//              babad, invalid
// midd=2.5,    ba, invalid

// T: O(n^2), S: O(1)