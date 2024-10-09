# 332. Reconstruct Itinerary

## Intuition
DFS
1. Form a graph.
2. Start by trying the paths sorted from smallest.
3. Skip duplicate destination.

## Approach

// Normal Case
// tickets=[["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
// graph=[JFK:[KUL, NRT], NRT:[JFK]]
// JFK, find [KUL, NRT],    pick KUL, path=[JFK, KUL], graph=[JFK:[NRT], NRT:[JFK]]
//      KUL, not find, restore graph=[JFK:[KUL, NRT], NRT:[JFK]]
//                          pick NRT, path=[JFK, NRT], graph=[JFK:[KUL], NRT:[JFK]]
//      NRT, find [JFK],    pick JFK, path=[JFK, NRT, JFK], graph=[JFK:[KUL]] >> when I try the next priority, how to restore the original graph
//          JFK, find [KUL],    pick KUL, path=[JFK, NRT, JFK, KUL], graph=[], return!

// Duplicate Tickets Case
// tickets=[["JFK","KUL"],["JFK","KUL"],["JFK","KUL"],["KUL","JFK"],["KUL","JFK"],["JFK","NRT"],["NRT","JFK"]]
// graph=[JFK:[KUL, KUL, KUL, NRT], NRT:[JFK], KUL:[JFK, JFK]]
// JFK, find [KUL, NRT],    pick KUL, path=[JFK, KUL], graph=[JFK:[KUL, KUL, NRT], NRT:[JFK], KUL:[JFK, JFK]]
//      KUL, not find, restore graph=[JFK:[KUL, KUL, KUL, NRT], NRT:[JFK], KUL:[JFK, JFK]]
//      The same process as the `Normal Case`, but since we already try `KUL`, we can skip it now.

Total time complexity = O(m) + [O(mlogm) or O(m) or O(n)] + O(m)
    if m >> n, O(mlogm)
    if n >> m, O(n)
Space complexity = O(m), because we store m tickets data into the graph.