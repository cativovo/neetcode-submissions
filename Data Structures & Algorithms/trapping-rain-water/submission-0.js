class Solution {
    /**
     * @param {number[]} height
     * @return {number}
     */
    trap(height) {
      const maxLefts = new Array(height.length);
      let maxL = 0;
      for (let i = 0; i < height.length; i++) {
        maxL = Math.max(maxL, height[i - 1] ?? 0)
        maxLefts[i] = maxL
      }

      const maxRights = new Array(height.length);
      let maxR = 0;
      for (let i = height.length - 1; i >= 0; i--) {
        maxR = Math.max(maxR, height[i + 1] ?? 0);
        maxRights[i] = maxR;
      }

      let result = 0;
      for (let i = 0; i < height.length; i++) {
        const area = Math.min(maxLefts[i], maxRights[i]) - height[i];
        if (area <= 0) {
            continue;
        }
        result += area;
      }

      return result;
    }
}
