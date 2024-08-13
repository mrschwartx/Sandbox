/**
 * @param {string[]} names
 * @param {number[]} heights
 * @return {string[]}
 */
var sortPeople = function (names, heights) {
  // Combine names and heights into an array of objects
  let people = names.map((name, index) => ({
    name: name,
    height: heights[index],
  }));

  // Sort the array of objects by height in descending order
  people.sort((a, b) => b.height - a.height);

  // Extract the names from the sorted array
  let sortedNames = people.map((person) => person.name);

  return sortedNames;
};
