class Solution {
  /**
   * @param {number[]} nums
   * @return {boolean}
   */
  hasDuplicate(nums) {
    const m = {};

    for (let i = 0; i < nums.length; i++) {
      const v = nums[i];
      if (m[v]) {
        return true;
      }

      m[v] = true;
    }

    return false;
  }
}