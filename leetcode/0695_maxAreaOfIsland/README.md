# 695. Max Area of Island

## Intuition
DFS to search all lands.

## Approach
// step 1. iterate all cell
// step 2. start dfs if the cell is land
// step 3. dfs all lands within the group and return area while mark them as water.
// step 4. update max area.

// T:O(n^2) // because we will search all lands.