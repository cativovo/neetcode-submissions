class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
        let result = 0;

        for (let l = 0; l < heights.length; l++) {
            for (let r = l + 1; r < heights.length; r++) {
                const area = (r - l) * Math.min(heights[l], heights[r])
                result = Math.max(result, area)
            }
        }

            return result;

    }

}
