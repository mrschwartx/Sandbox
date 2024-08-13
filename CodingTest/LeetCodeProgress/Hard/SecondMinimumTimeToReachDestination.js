/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number} time
 * @param {number} change
 * @return {number}
 */
function secondMinimum(n, edges, time, change) {
  // Create the graph using an adjacency list
  const graph = Array.from({ length: n + 1 }, () => []);
  const q = [];
  q.push([1, 0]);

  // minTime[i][0] := the first minimum time to reach the node i
  // minTime[i][1] := the second minimum time to reach the node i
  const minTime = Array.from({ length: n + 1 }, () => [Infinity, Infinity]);
  minTime[1][0] = 0;

  // Fill the adjacency list based on the given edges
  for (const [u, v] of edges) {
    graph[u].push(v);
    graph[v].push(u);
  }

  // BFS with time adjustment for traffic signals
  while (q.length > 0) {
    const [i, prevTime] = q.shift();
    const numChangeSignal = Math.floor(prevTime / change);
    const waitTime =
      numChangeSignal % 2 === 0 ? 0 : change - (prevTime % change);
    const newTime = prevTime + waitTime + time;

    for (const j of graph[i]) {
      if (newTime < minTime[j][0]) {
        minTime[j][0] = newTime;
        q.push([j, newTime]);
      } else if (minTime[j][0] < newTime && newTime < minTime[j][1]) {
        if (j === n) {
          return newTime;
        }
        minTime[j][1] = newTime;
        q.push([j, newTime]);
      }
    }
  }

  return -1; // If no second minimum time is found (though the problem guarantees one).
}
