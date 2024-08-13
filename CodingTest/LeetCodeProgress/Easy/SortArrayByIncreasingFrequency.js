/**
 * @param {number[]} nums
 * @return {number[]}
 */
var frequencySort = function (nums) {
  // Count frequencies
  const frequencyMap = new Map();
  nums.forEach((num) => {
    frequencyMap.set(num, (frequencyMap.get(num) || 0) + 1);
  });

  // Convert map to array of tuples and sort
  const sortedEntries = Array.from(frequencyMap.entries()).sort(
    ([num1, freq1], [num2, freq2]) => {
      return freq1 - freq2 || num2 - num1;
    }
  );

  // Reconstruct the sorted array
  const result = [];
  sortedEntries.forEach(([num, freq]) => {
    result.push(...Array(freq).fill(num));
  });

  return result;
};
