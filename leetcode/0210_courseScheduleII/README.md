# 210. Course Schedule II

## Intuition
Graph, BFS.

## Approach
Imaging it's a graph, there will have one or more root courses without any prerequisites courses.
Then we keep check next courses after finishing current courses.

// num=4, prere=[[1,0],[2,0],[3,1],[3,2]]
// childMap=[0:[1,2], 1:[3], 2:[3]]
// countParents=[0, 1, 1, 2]
// init queue=[0]
// round 1, ans=[0], queue=[1,2]
// round 2, ans=[1,2], queue=[3]
// round 3, ans=[1,2,3], queue=[]
// return ans=[1,2,3]