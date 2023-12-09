# Intuition
It a directed graph problem, so my intuition is using DFS or BFS.

# Approach
1. Convert the prerequisities to a map[int][]int, then I can remember all the prerequired course numbers.
2. For each all courses to find if the course has cycle.
3. Implement the DFS function to check if the course has cycle.

The DFS function need a `visitedList` to store if the course has visited.
When I walked through the course, `visitedList[course] = true`.
But I got a problem:
```
Input: 5, [[0,1],[0,2],[0,3],[1,3]]

The relationship will look like:
0 is blocked by [1,2,3]
1 is blocked by [3]

But the executing order is ...
Stack 01
start: 0
visited=[0:1]
prereq course: 1,2,3

Stack 02
start: 1
visited=[0:1,1:1]
prereq course: 3

Stack 03
start: 3
visited=[0:1,1:1,3:1]
prereq course: nil
It's end, back to `Stack 02`

It's end, back to `Stack 01`

Stack 01
start: 2
visited=[0:1,1:1,3:1,2:1]
prereq course: nil

Stack 01
start: 3
visited=[0:1,1:1,3:1,2:1], is visited, return false
```

But accroding the solution of [ayoubzulfiqar](https://leetcode.com/problems/course-schedule/solutions/3759096/go-fast-simple-fast-and-easy-with-explanation/).
He changes the visited value to be `-1` at the end of the current stack:

```
...

Stack 03
start: 3
visited=[0:1,1:1,3:1]
prereq course: nil
It's end, visited=[0:1,1:1,3:-1], back to `Stack 02`.

...

Stack 01
start: 3
visited=[0:1,1:1,3:-1,2:1], is already checked, return true
```


# Complexity
- Time complexity:
N is numCourse.
M is the number of prerequisites.
O(N+M)

- Space complexity:
Graph: O(N+M).
Visited Array: O(N).
Total: Graph: O(N+M).