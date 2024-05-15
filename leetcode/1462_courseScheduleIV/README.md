# 1462. Course Schedule IV

## Intuition
Using DFS and cache to calculate all prerequisites of the courses.

## Approach
// n=numCourses
// len(prerequisites)=n^2
// len(queries)=m

// convert prerequisites to graph=map[course][]prerequistes >> O(n^2)
// DFS to find all prerequesites of the courses, prereqMap=map[course][]prerequistes >> O(n)
// check if the prerequisite is exist from queries >> O(m*n)

// Total:O(n^2+n+m*n) >> O(n^2+m*n)