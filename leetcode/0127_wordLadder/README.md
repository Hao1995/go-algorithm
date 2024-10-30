# 127. Word Ladder

## Intuition
// n: wordList size
// m: word lenth

// ["hot","dot","dog","lot","log","cog"]
// build a graph >> O(n^2)
//   check two strings are diff by one letter >> O(m)
// >> O(n^2 * m)
// [hot:[dot, lot], dot:[hot, lot, dog], dog:[dot, log, cog], lot:[hot, dot, log], cog:[dog, log]]
// 
// Since the n is 5000, m is 10. So the above time complexity would be 25*10^7. It must exceed the time limitation.

## Approach
Ref: https://youtu.be/h9iTnkgv05E?si=xve8-cwmXzu2oELu

// Store the graph with wildcard pattern. For example, the word is "hot", then we will store "\*ot", "h\*t", and "ho\*"
// build a graph >> O(n)
//   convert to wildcard patterns and insert to neighbors >> O(m)
// >> O(n * m)

// BFS: Find the shortest path from beginWord to endWord.
// Even though we won't take the path we've already traveled, some related paths will still remain in the queue. >> The worst case of the total paths is O(n^2).

// Detail explanation: From `dog`, `log`, and `aog` to `cog`, we only insert `cog` once because we have a `cache` map. But we still need to pop all `*og` words to check the next path.
// 
// layer:0      hot
// layer:1 dot  lot  aot
// layer:2 dog, log, aog
// layer:3      cog
// layer:4      cag