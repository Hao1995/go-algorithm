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

// Dijkstra O((V+E)logV)
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