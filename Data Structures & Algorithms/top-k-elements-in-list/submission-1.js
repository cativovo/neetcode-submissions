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

    const a = Array.from({ length: nums.length + 1 }, () => []);
    for (const k in m) {
      const v = m[k];
      a[v].push(parseInt(k));
    }

    const r = [];
    for (let i = a.length - 1; i >= 0; i--) {
      r.push(...a[i]);

      if (r.length === k) {
        return r;
      }
    }

    return [];
 

    }
}
