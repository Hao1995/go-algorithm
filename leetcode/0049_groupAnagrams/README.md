# 49. Group Anagrams

## Intuition

## Approach
Every string in the input are at most 100 length.
The length of input array is between 1 to 10^4.
If we foreach each string then sort it, the worse case is 10^4 * 100 = 10^6.
The time complexity is O(n * mlogm).

The k is the unique string in the input array.
So the space complexity is O(k * n).

// ["eat","tea","tan","ate","nat","bat"]
// "eat" >> ["aet"]=["eat"]
// "tea" >> ["aet"]=["eat","tea"]
// "tan" >> ["ant"]=["tan"]
// "ate" >> ["aet"]=["eat","tea","ate"]
// "nat" >> ["ant"]=["tan","nat"]
// "bat" >> ["abt"]=["bat"]
// ans >> [["bat"],["eat","tea","ate"],["tan","nat"]]