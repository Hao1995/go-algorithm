# 1786. Number of Restricted Paths From First to Last Node

## Intuition
Dijkstra + Dynamic Programming

## Approach
// 1 --- 3 -
// | \   |  |
// |   - 2  |
// |     |  |
// 4 --- 5 -

// Convert edges to graph >> O(E)
// Create cache >> O(V)

// Dijkstra O((V+E)logV) [1]
// start from 5
// choose [5,0], cache=[x,k,k,k,k,0], neig=[[3,1],[2,2],[4,10]], queue=[[3,1],[2,2],[4,10]]
// choose [3,1], cache=[x,k,k,1,k,0], neig=[[1,3],[2,1],[5,1]], skip which in cache ans lte=[5,1], queue=[[2,2],[2,1+1=2],[1,3+1=4],[4,10]]
// choose [2,2], cache=[x,k,2,1,k,0], neig=[[3,1],[1,3],[5,2]], skip which in cache ans lte=[3,1],[5,2], queue=[[2,2],[1,4],[1,3+2=5],[4,10]]
// choose [2,2], in cache, pass, queue=[[1,4],[1,5],[4,10]]
// choose [1,4], cache=[x,4,2,1,k,0], neig=[[3,3],[4,2],[2,3]], skip which in cache ans lte=[3,3],[2,3], queue=[[1,5],[4,6],[4,10]]
// choose [1,5], in cache, pass, queue=[[4,6],[4,10]]
// choose [4,6], cache=[x,4,2,1,6,0], neig=[[1,2],[5,10]], skip ... [[1,2],[5,10]], queue=[[4,10]]
// choose [4,10], in cache, pass...
// cache=[x,4,2,1,6,0]

// Find valid number of paths >> O(V)
// start from 1, neig=[3,2,4], only [3,2] is valid, cache=[1,1,1,x,x]
//      start from j


[1]
There are too many versions of Dijkstra time complexity.
Examples(only discuss the main loop which extract from min heap):
[GeeksForGeeks](https://www.geeksforgeeks.org/time-and-space-complexity-of-dijkstras-algorithm/) >> O((V+E)logV)
[Youtube - Niema Moshiri](https://youtu.be/YMyO-yZMQ6g?si=6OiZ_xaG5T_cy13u) >> O(ElogE)

I asked ChatGPT to get a better answer.
```
    // Main loop to process vertices
    while minHeap is not empty:
        // Extract the vertex with minimum distance value >> O(logV)
        u = extractMin(minHeap)

        if u.dist >= dist[u]
            continue
        dist[u] = u.dist

        // Update distance value of u's neighbors >> O(E)
        for each neighbor v of u:
            // If the new distance is smaller than the current distance,
            // update the distance and decrease the key of the vertex in the minHeap
            alt = dist[u] + weight(u, v)
            if alt < dist[v]:
                dist[v] = alt
                decreaseKey(minHeap, v, alt) // >> O(logV)
```
It said the main loop is O((V + E)logV) where E is the number of edges. The loop runs V times (once for each vertex), and in each iteration, we perform extractMin operation on the minHeap, which costs O(logV), and potentially update the distance and decrease the key for each edge in the graph, which can be O(E) in total.

Why the main loop totally run V times not E times? >>  This is because Dijkstra's algorithm processes each vertex exactly once to find the shortest path from the source vertex to all other vertices.

It because we already sett a distance cache to avoid too many operations.

You may ask `There has a loop cost O(E) within the main loop of the min heap, why our time complexity isn't O(V) * O(E)?`.
It because we have a condition `if alt < dist[v]`, it would not insert the vertex again otherwise the short path comes up.