# 130. Surrounded Regions

## Intuition
Convert all edge region cells to visited status, then convert the unvisited 'O' to 'X'.

## Approach
1. Convert all edge region cells to visited.
2. Convert all unvisited 'O' to 'X'

Time Complexity: O(n*m)
Space Complexity: O(n*m) // visited matrix.