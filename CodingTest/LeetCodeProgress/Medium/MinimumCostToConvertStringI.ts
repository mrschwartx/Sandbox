function minimumCost(
  source: string,
  target: string,
  original: string[],
  changed: string[],
  cost: number[]
): number {
  // Define the graph
  const graph: Map<string, [string, number][]> = new Map();

  for (let i = 0; i < original.length; i++) {
    if (!graph.has(original[i])) {
      graph.set(original[i], []);
    }
    graph.get(original[i])!.push([changed[i], cost[i]]);
  }

  // Dijkstra's algorithm to find shortest path from a start character
  function dijkstra(start: string): Map<string, number> {
    const minCost: Map<string, number> = new Map();
    const pq: [number, string][] = [];
    for (let i = 0; i < 26; i++) {
      minCost.set(String.fromCharCode(97 + i), Infinity);
    }
    minCost.set(start, 0);
    pq.push([0, start]);

    while (pq.length > 0) {
      pq.sort((a, b) => a[0] - b[0]);
      const [currentCost, u] = pq.shift()!;
      if (currentCost > minCost.get(u)!) continue;
      if (graph.has(u)) {
        for (const [v, weight] of graph.get(u)!) {
          const cost = currentCost + weight;
          if (cost < minCost.get(v)!) {
            minCost.set(v, cost);
            pq.push([cost, v]);
          }
        }
      }
    }
    return minCost;
  }

  // Calculate the minimum costs for all transformations
  const minCosts: Map<string, Map<string, number>> = new Map();
  const uniqueChars: Set<string> = new Set([...source, ...target]);

  for (const char of uniqueChars) {
    minCosts.set(char, dijkstra(char));
  }

  // Calculate total cost to transform source to target
  let totalCost = 0;
  for (let i = 0; i < source.length; i++) {
    if (source[i] === target[i]) continue;
    const cost = minCosts.get(source[i])!.get(target[i])!;
    if (cost === Infinity) return -1;
    totalCost += cost;
  }

  return totalCost;
}
