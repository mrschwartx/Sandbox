/**
 * @param {number[]} mapping
 * @param {number[]} nums
 * @return {number[]}
 */
var sortJumbled = function (mapping, nums) {
  function mapValue(num) {
    return parseInt(
      num
        .toString()
        .split("")
        .map((digit) => mapping[digit])
        .join(""),
      10
    );
  }

  const mappedNums = nums.map((num) => [num, mapValue(num)]);
  mappedNums.sort((a, b) => a[1] - b[1]);

  return mappedNums.map((pair) => pair[0]);
};
