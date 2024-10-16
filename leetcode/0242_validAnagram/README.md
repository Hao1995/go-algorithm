# 242. Valid Anagram

## Intuition
Hash Map.
1. Count each character
2. Compare both of hash map are same

## Approach
// s="anagram"
// smap=[a:3,n:1,g:1,r:1,m:1]
// t="nagaram"
// tmap=[n:1,a:3,g:1,r:1,m:1]
// iterate smap
// a:3, tmap[a]=3
// n:1, tmap[n]=1
// g:1, tmap[g]=1
// r:1, tmap[r]=1
// m:1, tmap[m]=1

// T:O(n)
// S:O(n)
