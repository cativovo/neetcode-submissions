class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
           if (nums.length === 0) {
      return 0;
    }

    const o = {};

    for (const v of nums) {
      o[v] = true;
    }

    let result = 1;

    for (let v of nums) {
      const isStart = !o[v - 1];
      let count = 1;
      while (isStart) {
        v++;
        if (o[v]) {
          count++;
          continue;
        }

        if (count > result) {
          result = count;
        }

        break;
      }
    }

    return result;

    }
}
