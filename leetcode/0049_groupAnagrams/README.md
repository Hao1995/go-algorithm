# 49. Group Anagrams

## Intuition

## Approach
1. Sorting + HashMap
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

2. HashMap + Letter Array
// HashMap
// strs=["eat","tea","tan","ate","nat","bat"]
// "eat", [1,0,0,0,1,...1,...,0] -> "1001..1..0" O(n), push to map
// ...iterate k str, total would be O(kn)