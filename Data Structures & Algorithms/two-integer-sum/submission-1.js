class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {

            const d = nums.reduce((acc, v, i) => {
      acc[target - v] = i;
      return acc;
    }, {});

    for (let i = 0; i < nums.length; i++) {
      const v = d[nums[i]];
      if (v !== undefined && v !== i) {
        return [v, i];
      }
    }

    return [-1, -1];

    }
}
