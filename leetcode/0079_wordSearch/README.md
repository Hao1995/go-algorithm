# 79. Word Search

## Intuition
DFS!
We need to search all direction to check if the character from the `board` is matched the character from the `word`.

## Approach
* For loop each point on the `board` to check if the `word` existed
* Start to traversal possible routes and mark the possible route to be visited

But the problem I met is I forgot to mark the failed rout to be unvisited.
```go
// I mark the current point to be `visited`
if visited[r][c] {
    return false
}
visited[r][c] = true

// Then I start traversal each direction
var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
for _, dir := range dirs {
    if dfs(board, visited, []int{r + dir[0], c + dir[1]}, word) {
        return true
    }
}
```

The above solution will make the following problem:

```
// [Input]
// board=
// A, B, C, E
// S, F, E, S
// A, D, E, E
// word=ABCEFSADEESE

// step 001. Start from A(0,0), yes, mark A is visited, word=BCEFSADEESE
// step 002. A right, B yes, mark B is visited, word=CEFSADEESE
// step 003. B right, C yes, mark C is visited, word=EFSADEESE
// step 004. C right, E yes, mark E is visited, word=FSADEESE
// step 005. E right, out of boundary, next
// step 006. E down, S no, next
// step 007. E left, C is visited, next
// step 008. E top, out of boundary, next
// step 009. C down, E yes, mark E is visited, word=FSADEESE
// step 010. E right, S no, break
// step 011. E down, E no, break
// step 012. E left, F yes, makk F is visited, word=SADEESE
// step 013. F right, E is visited, next
// step 014. F down, D no, next
// step 015. F left, S yes, mark S is visited, word=ADEESE
// step 016. S right, F is visited, next
// step 017. S down, A yes, mark A is visited, word=DEESE
// step 018. A right, D yes, mark D is visited, word=EESE
// step 019. D right, E yes, mark E is visited, word=ESE
// step 021. E right, E yes, mark E is visited, word=SE
// step 022. E right, out of boundary, next
// step 023. E down, out of boundary, next
// step 024. E left, E is visited, next
// step 025. E up, S ues, makr S is visited, word=E
// step 026. S right, out of boundary
// step 027. S down, E is visited, next
// step 028. S left, E is visited, next
// step 029. S top, E is visited, next

That's why this solution will failed.
Because the path from C(0,2) to E(0,3) is failed.
Then we try the new direction which from C(0,2) to E(1,2) at `step 009`
```

We need to mark the previous wrong path to be unvisited. (mark E(0,3) is unvisited)
```go
// ...

// Then I start traversal each direction
var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
for _, dir := range dirs {
    if dfs(board, visited, []int{r + dir[0], c + dir[1]}, word) {
        return true
    }
}

visited[r][c] = true
return false
```
Then it's work!
```
// step 008-2. C right, E path failed, mark E is unvisited, word=FSADEESE
// ...
// step 029. S top, E yes, mark E is visited, word=
// step 030. E right, length of `word` is empty, return true
```
