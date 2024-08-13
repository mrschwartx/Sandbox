/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArray = function (nums) {
  const mergeSort = (A, l, r) => {
    if (l >= r) return;

    const merge = (A, l, m, r) => {
      const sorted = new Array(r - l + 1);
      let k = 0;
      let i = l;
      let j = m + 1;

      while (i <= m && j <= r) {
        if (A[i] < A[j]) {
          sorted[k++] = A[i++];
        } else {
          sorted[k++] = A[j++];
        }
      }

      while (i <= m) {
        sorted[k++] = A[i++];
      }

      while (j <= r) {
        sorted[k++] = A[j++];
      }

      for (let p = 0; p < sorted.length; p++) {
        A[l + p] = sorted[p];
      }
    };

    const m = Math.floor((l + r) / 2);
    mergeSort(A, l, m);
    mergeSort(A, m + 1, r);
    merge(A, l, m, r);
  };
  mergeSort(nums, 0, nums.length - 1);
  return nums;
};
