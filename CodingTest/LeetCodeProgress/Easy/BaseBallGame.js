/**
 * @param {string[]} operations
 * @return {number}
 */
var calPoints = function (operations) {
  let scores = [];

  for (let ops of operations) {
    if (ops == "+") {
      scores.push(scores[scores.length - 1] + scores[scores.length - 2]);
    } else if (ops == "D") {
      scores.push(scores[scores.length - 1] * 2);
    } else if (ops == "C") {
      scores.pop();
    } else {
      scores.push(parseInt(ops));
    }
  }

  return scores.reduce((a, b) => a + b, 0);
};
