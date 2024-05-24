# 212. Word Search II

## Intuition
DFS + cache >> burst force
// burst force
// 1. for each words >> len(words) >> O(k)
// 2. start from each point in board >> O(mn)
// 3. check each letter >> len(words[i]) >> min (O(v), O(mn))
// >> O(k*(mn)^2)

## Approach
Trie + DFS + cache
[NeetCode](https://youtu.be/asbcE9mZz_U?si=UGI3N_SFsdJKECOf)

// trie
// 1. build a trie from words >> T:O(kv) S:O(kv)
// 2. start from each point on the board >> O(mn)
// 3. DFS through all adjacent cells to check if it exist in the Trie >> O(mn)
// >> O(kv+(mn)^2)