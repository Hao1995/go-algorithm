class Solution {
    /**
     * @param {number} n
     * @param {number[][]} edges
     * @returns {number}
     */
    countComponents(n, edges) {
        // Create a graph from edges >> O(e)
        // hash.keys.length
        const graph = {};
        for (const edge of edges) {
            if (!graph[edge[0]]) graph[edge[0]] = [];
            if (!graph[edge[1]]) graph[edge[1]] = [];
            graph[edge[0]].push(edge[1]);
            graph[edge[1]].push(edge[0]);
        }

        // Array to track visited nodes >> O(n)
        const visited = new Array(n).fill(false);

        // Helper function to perform DFS >> O(n)
        const dfs = (start) => {
            if (visited[start]) {
                return;
            }
            visited[start] = true;

            const neighbors = graph[start] || [];
            for (const neighbor of neighbors) {
                dfs(neighbor);
            }
        }

        // Count components >> O(n)
        let ans = 0;
        for (let i = 0; i < n; i++) {
            if (!visited[i]) {
                dfs(i);
                ans++;
            }
        }

        return ans;
    }
}