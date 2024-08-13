function findTheCity(
  n: number,
  edges: number[][],
  distanceThreshold: number
): number {
  // Define a large number to represent infinity
  const INF = Number.MAX_VALUE;

  // Initialize the distance matrix
  const dist: number[][] = Array.from({ length: n }, () => Array(n).fill(INF));

  // Set distance to self as 0
  for (let i = 0; i < n; i++) {
    dist[i][i] = 0;
  }

  // Set the initial distances based on the edges
  for (const [u, v, w] of edges) {
    dist[u][v] = w;
    dist[v][u] = w;
  }

  // Floyd-Warshall algorithm to find all-pairs shortest paths
  for (let k = 0; k < n; k++) {
    for (let i = 0; i < n; i++) {
      for (let j = 0; j < n; j++) {
        if (dist[i][k] < INF && dist[k][j] < INF) {
          dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);
        }
      }
    }
  }

  // Find the city with the smallest number of reachable cities within the distance threshold
  let minCount = INF;
  let bestCity = -1;

  for (let i = 0; i < n; i++) {
    let count = 0;
    for (let j = 0; j < n; j++) {
      if (dist[i][j] <= distanceThreshold) {
        count++;
      }
    }

    // Update the city with the smallest count, or the greatest number if tied
    if (count < minCount || (count === minCount && i > bestCity)) {
      minCount = count;
      bestCity = i;
    }
  }

  return bestCity;
}
