class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
            const m = nums.reduce((acc, v) => {
      acc[v] = (acc[v] ?? 0) + 1;
      return acc;
    }, {});

    return Object.entries(m)
      .sort((a, b) => b[1] - a[1])
      .map((v) => parseInt(v[0]))
      .slice(0, k);

    }
}
