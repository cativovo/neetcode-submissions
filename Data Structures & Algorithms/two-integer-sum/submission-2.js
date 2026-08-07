class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
    const d = {};

    for (let i = 0; i < nums.length; i++) {
      const num = nums[i];
      const f = d[num];
      if (f !== undefined) {
        return [f, i];
      }

      d[target - num] = i;
    }

    return [-1, -1];
    }
}
