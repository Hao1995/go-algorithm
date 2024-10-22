class Solution {
    /**
     * @param {number} n
     * @param {number[][]} edges
     * @returns {boolean}
     */
    validTree(n, edges) {
        if (n === 1) return true;
        if (edges.length !== n - 1) return false;

        const graph = new Map();
        edges.forEach(([u, v]) => {
            graph.set(u, (graph.get(u) || []).concat(v));
            graph.set(v, (graph.get(v) || []).concat(u));
        });

        const visited = new Set();

        function dfs(prev, curr) {
            if (visited.has(curr)) return false;

            visited.add(curr);

            for (const neighbor of graph.get(curr) || []) {
                if (neighbor !== prev && !dfs(curr, neighbor)) {
                    return false;
                }
            }
            return true;
        }

        if (!dfs(-1, 0)) {
            return false;
        }

        return visited.size === n;
    }
}
