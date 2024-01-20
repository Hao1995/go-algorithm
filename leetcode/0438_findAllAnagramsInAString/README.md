# 438. Find All Anagrams in a String

## Intuition
1. Two pointer
2. Sliding window

## Approach
// s="cbba"
// p="cbaa"

// s="cbaebabacd"
// p="abc" => map[byte]int32=["a":1, "b":1, "c":1]
// c, map[byte]int32=["a":1, "b":1, "c":0], start=0
// b, map[byte]int32=["a":1, "b":0, "c":0]
// a, map[byte]int32=["a":0, "b":0, "c":0]
// e, not found, isAnagram=true, ans=[0], initMap, start=-1
// b, map[byte]int32=["a":1, "b":0, "c":1], start=4
// a, map[byte]int32=["a":0, "b":0, "c":1]
// b, count <= 0, isAnagram=false, initMap, start=-1, i--
// *b, map[byte]int32=["a":1, "b":0, "c":1], start=6
// a, map[byte]int32=["a":0, "b":0, "c":1]
// c, map[byte]int32=["a":0, "b":0, "c":0]
// d, not found, isAnagram=true, ans=[0,6], initMap, start=9

// s="abab", p="ab", pMap=[a:1,b:1], start=-1
// a, [a:0,b:1], start=0
// b, [a:0,b:0]
// a, count<=0, isAnagram=true, ans=[0], addStartBack=[a:1,b:0]