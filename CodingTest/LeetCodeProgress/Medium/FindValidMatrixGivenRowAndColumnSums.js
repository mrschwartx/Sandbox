/**
 * @param {number[]} rowSum
 * @param {number[]} colSum
 * @return {number[][]}
 */
var restoreMatrix = function (rowSum, colSum) {
  const rows = rowSum.length;
  const cols = colSum.length;
  const matrix = Array.from({ length: rows }, () => Array(cols).fill(0));

  let i = 0;
  let j = 0;

  while (i < rows && j < cols) {
    const minVal = Math.min(rowSum[i], colSum[j]);
    matrix[i][j] = minVal;
    rowSum[i] -= minVal;
    colSum[j] -= minVal;

    if (rowSum[i] === 0) i++;
    if (colSum[j] === 0) j++;
  }

  return matrix;
};
